package mysql

import (
	"database/sql"
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
