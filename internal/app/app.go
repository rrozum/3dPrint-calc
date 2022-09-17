package app

import (
	"3dPrintCalc/internal/config"
	"3dPrintCalc/pkg/database/sqlite3"
	"3dPrintCalc/view/desktop"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

const (
	desktopMode config.AppMode = "desktop"
)

var (
	name sql.NullString
)

type testTable struct {
	id    int
	name  string
	price int
}

func Run(configPath string) {
	cfg, err := config.Init(configPath)
	if err != nil {
		fmt.Println(err)

		return
	}

	dbConnection, err := sqlite3.NewConnection(cfg.Database.Sqlite3.LocalPath)

	rows, err := dbConnection.Query("select * from test_table")

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)

	var testRows []testTable

	for rows.Next() {
		testRow := testTable{}

		err := rows.Scan(&testRow.id, &testRow.name, &testRow.price)
		if err != nil {
			continue
		}

		testRows = append(testRows, testRow)
	}

	switch cfg.AppMode {
	case desktopMode:
		desktop.Run(testRows[0].name)
	default:
		fmt.Println("app_mode not set!")
	}
}
