package repository

import (
	"3dPrintCalc/internal/domain"
	"database/sql"
)

type ApplicationSettingsRepo struct {
	db *sql.DB
}

func NewApplicationSettingsRepo(db *sql.DB) *ApplicationSettingsRepo {
	return &ApplicationSettingsRepo{
		db: db,
	}
}

func (r *ApplicationSettingsRepo) Save(settings domain.ApplicationSetting) error {
	_, err := r.db.Exec(
		"insert into application_settings (key, value) values ($1, $2)",
		settings.Key,
		settings.Value,
	)

	return err
}

// GetAll TODO добавить пагинацию
func (r *ApplicationSettingsRepo) GetAll() ([]domain.ApplicationSetting, error) {
	rows, err := r.db.Query("select * from application_settings")

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)

	var applicationSettings []domain.ApplicationSetting

	for rows.Next() {
		applicationSettingRow := domain.ApplicationSetting{}

		err := rows.Scan(&applicationSettingRow.Key, &applicationSettingRow.Value)
		if err != nil {
			continue
		}

		applicationSettings = append(applicationSettings, applicationSettingRow)
	}

	return applicationSettings, err
}
