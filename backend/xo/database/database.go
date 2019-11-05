package database

import (
	"database/sql"
	"github.com/xo/dburl"

	_ "github.com/go-sql-driver/mysql"
)

func New(dataSourceName string) (*sql.DB, error) {
	db, err := dburl.Open(dataSourceName)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}