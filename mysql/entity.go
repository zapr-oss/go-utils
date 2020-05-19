package mysql

import (
	"database/sql"
	"log"
	"github.com/zapr-oss/go-utils/graphite"
	"github.com/zapr-oss/go-utils/mysql/config"
	"github.com/zapr-oss/go-utils/mysql/mysqlutil"
	"github.com/zapr-oss/go-utils/stringutil"
	"fmt"
	"time"
)

// Connect DB and return mysql entity
func Connect(config mysqlconfig.Config) (*Entity, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.User, config.Password,
		config.Host, config.Port, config.Database))
	if err != nil {
		graphite.GetCounter("MySQLConnectionError").Inc()
		log.Println("Error connecting to mysql, Error: ", err)
		return nil, err
	}

	return &Entity{
		DB:         db,
		RetryCount: config.RetryCount,
	}, nil
}

type Entity struct {
	DB         *sql.DB
	RetryCount int
}

// This function is used to insert multiple values in one query.
/*
@params:
	query: a query containing bindVar string
	bindVar: its basically used to create a query, eg. `Insert INTO abc (column1, column2, column3) VALUES $$`. `$$` is the bindVar here
	noOfParams: 3 in the above example. Its the number of columns to insert
	args: list of arguments for the query.
*/
func (e *Entity) InsertMultiple(query string, bindVar string, noOfParams int, args ...interface{}) (sql.Result, error) {

	query, err := mysqlutil.CreateMultipleQuery(query, bindVar, noOfParams, len(args))

	if err != nil {
		return nil, err
	}

	return e.DB.Exec(query, args...)
}

func (e *Entity) Exec(query string, args ...interface{}) (sql.Result, error) {
	return e.DB.Exec(query, args...)
}

func (e *Entity) Insert(query string, args ...interface{}) (sql.Result, error) {
	return e.DB.Exec(query, args...)
}

func (e *Entity) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return e.DB.Query(query, args...)
}

func (e *Entity) QueryRow(query string, args ...interface{}) *sql.Row {
	return e.DB.QueryRow(query, args...)
}

func (e *Entity) ExecuteQuery(tx *sql.Tx, query string, args ...interface{}) (sql.Result, error) {
	return e.executeQueryWithRetry(tx, query, 1, args...)
}

func (e *Entity) executeQueryWithRetry(tx *sql.Tx, query string, iterationCount int,
	args ...interface{}) (sql.Result, error) {
	var res sql.Result
	var err error
	if tx != nil {
		res, err = tx.Exec(query, args...)
	} else {
		res, err = e.DB.Exec(query, args...)
	}

	if err != nil {
		if stringutil.CaseInsensitiveContains(err.Error(), "deadlock") {
			if iterationCount < e.RetryCount { // Try again after a delay
				time.Sleep(10 * time.Second)
				return e.executeQueryWithRetry(tx, query, iterationCount+1, args...)
			}
		}
	}
	return res, err
}

func (e *Entity) StartTransaction() (*sql.Tx, error) {
	tx, err := e.DB.Begin()
	if err != nil {
		log.Println("Failed to start transaction on DB.", err)
		return nil, err
	}
	return tx, nil
}

func (e *Entity) CommitTransaction(tx *sql.Tx) error {
	err := tx.Commit()
	if err != nil {
		log.Println("Failed to commit transaction on DB.", err)
		return err
	}
	return nil
}

func (e *Entity) RollbackTransaction(tx *sql.Tx) error {
	err := tx.Rollback()
	if err != nil {
		log.Println("Failed to rollback transaction on DB.", err)
		return err
	}
	return nil
}

func (e *Entity) Close() error {
	return e.DB.Close()
}
