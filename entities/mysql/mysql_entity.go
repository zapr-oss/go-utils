package mysql

import (
	"database/sql"
)

type Entity struct {
	DB         *sql.DB
	RetryCount int
}
