package repository

import (
	"3dPrintCalc/internal/domain"
	"database/sql"
)

type ProjectsRepo struct {
	db *sql.DB
}

func NewProjectsRepo(db *sql.DB) *ProjectsRepo {
	return &ProjectsRepo{
		db: db,
	}
}

func (r *ProjectsRepo) Save(project domain.Project) error {
	_, err := r.db.Exec(
		"insert into projects ("+
			"name,"+
			"price,"+
			"plastic_type,"+
			"plastic_price_by_coil,"+
			"plastic_price_by_kg,"+
			"operator_salary_by_month,"+
			"operator_salary_by_hour,"+
			"electricity_price,"+
			"rent_price,"+
			"minimal_price,"+
			"discount"+
			")"+
			" values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)",
		project.Name,
		project.Price,
		project.PlasticType,
		project.PlasticPriceByCoil,
		project.PlasticPriceByKg,
		project.OperatorSalaryByMonth,
		project.OperatorSalaryByHour,
		project.ElectricityPrice,
		project.RentPrice,
		project.MinimalPrice,
		project.Discount,
	)

	return err
}

func (r *ProjectsRepo) GetAll() ([]domain.Project, error) {
	rows, err := r.db.Query("select * from projects")

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)

	var projects []domain.Project

	for rows.Next() {
		projectRow := domain.Project{}

		err := rows.Scan(
			&projectRow.Id,
			&projectRow.Name,
			&projectRow.Price,
			&projectRow.PlasticType,
			&projectRow.PlasticPriceByCoil,
			&projectRow.PlasticPriceByKg,
			&projectRow.OperatorSalaryByMonth,
			&projectRow.OperatorSalaryByHour,
			&projectRow.ElectricityPrice,
			&projectRow.RentPrice,
			&projectRow.MinimalPrice,
			&projectRow.Discount,
		)
		if err != nil {
			continue
		}

		projects = append(projects, projectRow)
	}

	return projects, err
}
