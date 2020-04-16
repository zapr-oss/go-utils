package mysql_utils

import (
	"bitbucket.org/zapr/go-utils/entities/mysql"
	graphite "bitbucket.org/zapr/graphite_go"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

/*
	Returns a MySQL Entity struct with a connection open to the MySQL DB.
	On error, the error is returned.
*/
func Connect(config mysql.Config) (*mysql.Entity, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.User, config.Password,
		config.Host, config.Port, config.Database))
	if err != nil {
		graphite.GetCounter("MySQLConnectionError").Inc()
		log.Println("Error connecting to mysql, Error: ", err)
		return nil, err
	}

	return &mysql.Entity{
		DB:         db,
		RetryCount: config.RetryCount,
	}, nil
}
