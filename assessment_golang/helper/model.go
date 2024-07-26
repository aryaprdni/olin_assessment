package helper

import (
	"aryaprdni/assessment_golang/model/domain"
	"aryaprdni/assessment_golang/model/web"
)

func ToInputsResponse(inputs domain.Inputs, results domain.Results) web.InputsReponse {
	return web.InputsReponse{
		Id:      inputs.Id,
		Nums:    inputs.Nums,
		Target:  inputs.Target,
		Results: results.Indices,
	}
}
