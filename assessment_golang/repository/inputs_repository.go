package repository

import (
	"aryaprdni/assessment_golang/model/domain"
	"context"
	"database/sql"
)

type InputsRepository interface {
	Save(ctx context.Context, tx *sql.Tx, inputs domain.Inputs) domain.Inputs
	SaveResults(ctx context.Context, tx *sql.Tx, results domain.Results) (domain.Results, error)
}
