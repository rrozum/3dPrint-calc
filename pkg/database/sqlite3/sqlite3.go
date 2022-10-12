package sqlite3

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"path/filepath"
)

func NewConnection(dataSourceName string) (*sql.DB, error) {
	dataSourceName, _ = filepath.Abs(dataSourceName)

	db, err := sql.Open("sqlite3", dataSourceName)

	return db, err
}
