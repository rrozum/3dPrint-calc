package app

import (
	"3dPrintCalc/internal/config"
	"3dPrintCalc/internal/repository"
	"3dPrintCalc/internal/service"
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

	repos := repository.NewRepositories(dbConnection)
	services := service.NewServices(service.Deps{
		Repos: repos,
	})

	switch cfg.AppMode {
	case desktopMode:
		desktop.Run(services)
	default:
		fmt.Println("app_mode not set!")
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(dbConnection)
}
