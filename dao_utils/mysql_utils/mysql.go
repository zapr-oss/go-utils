package mysql_utils

import (
	"bitbucket.org/zapr/go-utils/entities/mysql"
	graphite "bitbucket.org/zapr/graphite_go"
	"database/sql"
	"fmt"
	"log"
)

/*
	Returns a MySQL Entity struct with a connection open to the MySQL DB.
	On error, the program exits.
*/
func Connect(config mysql.Config) *mysql.Entity {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.User, config.Password,
		config.Host, config.Port, config.Database))
	if err != nil {
		graphite.GetCounter("MySQLConnectionError").Inc()
		log.Fatal("Error connecting to mysql, Error: ", err)
	}

	return &mysql.Entity{
		DB:         db,
		RetryCount: config.RetryCount,
	}
}
