package db

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

func NewMySQLStorage(cfg mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}
	return db, nil
}

func InitStorage(db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		return err
	}
	// log.Println("DB: Successfully connected!")
	return nil
}
