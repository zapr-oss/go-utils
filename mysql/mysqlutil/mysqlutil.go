package mysqlutil

import (
	"database/sql"
	"fmt"
	"log"
	"github.com/zapr-oss/go-utils/graphite"
	"github.com/zapr-oss/go-utils/mysql/config"
	"math"
	"strings"
)

func CreateConnection(config mysqlconfig.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.User, config.Password,
		config.Host, config.Port, config.Database))
	if err != nil {
		graphite.GetCounter("MySQLConnectionError").Inc()
		log.Println("Error connecting to mysql, Error: ", err)
		return nil, err
	}

	return db, nil
}

func StartTransaction(db *sql.DB) (*sql.Tx, error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func CommitTransaction(tx *sql.Tx) error {
	err := tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func RollbackTransaction(tx *sql.Tx) error {
	err := tx.Rollback()
	if err != nil {
		return err
	}
	return nil
}


//Used to create an insert query with multiple parameters.
func CreateMultipleQuery(query, bindVar string, noOfParams, argLen int) (string, error) {

	if !strings.Contains(query, bindVar) {
		return "", fmt.Errorf("couldn't find bindVar in the given insertQuery, \nquery: %v\nbindVariable: %v", query, bindVar)
	}

	if math.Mod(float64(argLen), float64(noOfParams)) != 0.0 {
		return "", fmt.Errorf(
			"total number args can't be divided by noOfParams, "+
				"\narg length: %v\nnoOfParams: %v", argLen, noOfParams)
	}

	paramSetStr := createParamString(noOfParams, argLen/noOfParams)

	query = strings.Replace(query, bindVar, paramSetStr, 1)

	return query, nil
}

func createParamString(noOfParams, totalParamSets int) string {
	singleParamStr := "("

	i := 0
	for i < noOfParams {
		singleParamStr += "?,"
		i++
	}

	singleParamStr = singleParamStr[:len(singleParamStr)-1]
	singleParamStr += ")"
	i = 0

	paramSetStr := ""
	for i < totalParamSets {
		paramSetStr += singleParamStr + ","
		i++
	}

	paramSetStr = paramSetStr[:len(paramSetStr)-1]

	return paramSetStr
}
