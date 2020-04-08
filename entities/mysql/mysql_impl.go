package mysql

import (
	"database/sql"
	"log"
	"strings"
	"time"
)

func (entity *Entity) ExecuteQuery(tx *sql.Tx, query string, args ...interface{}) (sql.Result, error) {
	return entity.executeQueryWithRetry(tx, query, 1, args...)
}

func (entity *Entity) executeQueryWithRetry(tx *sql.Tx, query string, iterationCount int,
	args ...interface{}) (sql.Result, error) {

	var res sql.Result
	var err error
	if tx != nil {
		res, err = tx.Exec(query, args...)
	} else {
		res, err = entity.DB.Exec(query, args...)
	}

	if err != nil {
		if strings.Contains(err.Error(), "Deadlock") {
			if iterationCount < entity.RetryCount { // Try again after a delay
				time.Sleep(10 * time.Second)
				return entity.executeQueryWithRetry(tx, query, iterationCount+1, args...)
			}
		}
	}
	return res, err
}

func (entity *Entity) StartTransaction() *sql.Tx {
	tx, err := entity.DB.Begin()
	if err != nil {
		log.Println("Failed to start transaction on DB.", err)
		return nil
	}
	return tx
}

func (entity *Entity) CommitTransaction(tx *sql.Tx) bool {
	err := tx.Commit()
	if err != nil {
		log.Println("Failed to commit transaction on DB.", err)
		return false
	}
	return true
}

func (entity *Entity) RollbackTransaction(tx *sql.Tx) bool {
	err := tx.Rollback()
	if err != nil {
		log.Println("Failed to rollback transaction on DB.", err)
		return false
	}
	return true
}
