package repository

import (
	"3dPrintCalc/internal/domain"
	"database/sql"
)

type DetailsRepo struct {
	db *sql.DB
}

func NewDetailsRepo(db *sql.DB) *DetailsRepo {
	return &DetailsRepo{
		db: db,
	}
}

func (r *DetailsRepo) Save(detail domain.Detail) error {
	_, err := r.db.Exec(
		"insert into details ("+
			"name, "+
			"margin_ratio, "+
			"plastic_consumption, "+
			"print_time, "+
			"data_preparation_by_operator_time, "+
			"post_processing_by_operator_time, "+
			"project_id"+
			") values ($1, $2, $3, $4, $5, $6, $7)",
		detail.Name,
		detail.MarginRatio,
		detail.PlasticConsumption,
		detail.PrintTime,
		detail.DataPreparationByOperatorTime,
		detail.PostProcessingByOperatorTime,
		detail.ProjectId,
	)

	return err
}

func (r *DetailsRepo) GetByProject(projectId int) ([]domain.Detail, error) {
	rows, err := r.db.Query("select * from details where project_id = $1", projectId)

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)

	var details []domain.Detail

	for rows.Next() {
		detailRow := domain.Detail{}

		err := rows.Scan(
			&detailRow.Id,
			&detailRow.Name,
			&detailRow.MarginRatio,
			&detailRow.PlasticConsumption,
			&detailRow.PrintTime,
			&detailRow.DataPreparationByOperatorTime,
			&detailRow.PostProcessingByOperatorTime,
			&detailRow.ProjectId,
		)

		if err != nil {
			continue
		}

		details = append(details, detailRow)
	}

	return details, err
}
