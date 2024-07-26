package repository

import (
	"aryaprdni/assessment_golang/helper"
	"aryaprdni/assessment_golang/model/domain"
	"context"
	"database/sql"
	"fmt"
	"strings"
)

type InputsRepositoryImpl struct {
}

func NewInputsRepository() InputsRepository {
	return &InputsRepositoryImpl{}
}

func sliceToPostgresArray(arr []int) string {
	if len(arr) == 0 {
		return "{}"
	}
	strs := make([]string, len(arr))
	for i, v := range arr {
		strs[i] = fmt.Sprintf("%d", v)
	}
	return fmt.Sprintf("{%s}", strings.Join(strs, ","))
}

func (repository *InputsRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, inputs domain.Inputs) domain.Inputs {
	numsArray := sliceToPostgresArray(inputs.Nums)

	SQL := "INSERT INTO inputs(nums, target) VALUES ($1, $2) RETURNING id"
	row := tx.QueryRowContext(ctx, SQL, numsArray, inputs.Target)

	var id int
	err := row.Scan(&id)
	helper.PanicIfError(err)

	inputs.Id = id
	return inputs
}

func (repository *InputsRepositoryImpl) SaveResults(ctx context.Context, tx *sql.Tx, results domain.Results) (domain.Results, error) {
	indicesArray := sliceToPostgresArray(results.Indices)

	SQL := "INSERT INTO results(input_id, indices) VALUES ($1, $2) RETURNING id"
	row := tx.QueryRowContext(ctx, SQL, results.InputId, indicesArray)

	var id int
	err := row.Scan(&id)
	if err != nil {
		return results, err
	}

	results.Id = id
	return results, nil
}
