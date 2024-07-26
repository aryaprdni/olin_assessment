package service

import (
	"aryaprdni/assessment_golang/helper"
	"aryaprdni/assessment_golang/model/domain"
	"aryaprdni/assessment_golang/model/web"
	"aryaprdni/assessment_golang/repository"
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
)

type InputsServiceImpl struct {
	InputsRepository repository.InputsRepository
	DB               *sql.DB
	Validate         *validator.Validate
}

func NewInputsService(inputsRepository repository.InputsRepository, DB *sql.DB, validate *validator.Validate) InputsService {
	return &InputsServiceImpl{
		InputsRepository: inputsRepository,
		DB:               DB,
		Validate:         validate,
	}
}

func (service *InputsServiceImpl) Create(ctx context.Context, request web.InputsCreateRequest) web.InputsReponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	inputs := domain.Inputs{
		Nums:   request.Nums,
		Target: request.Target,
	}

	inputs = service.InputsRepository.Save(ctx, tx, inputs)

	results := findResults(inputs)

	results.InputId = inputs.Id

	results, err = service.InputsRepository.SaveResults(ctx, tx, results)
	helper.PanicIfError(err)

	return helper.ToInputsResponse(inputs, results)
}

func findResults(inputs domain.Inputs) domain.Results {
	nums := inputs.Nums
	target := inputs.Target
	indices := []int{}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				indices = append(indices, i, j)
			}
		}
	}

	return domain.Results{
		Indices: indices,
	}
}
