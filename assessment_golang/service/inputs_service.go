package service

import (
	"aryaprdni/assessment_golang/model/web"
	"context"
)

type InputsService interface {
	Create(ctx context.Context, request web.InputsCreateRequest) web.InputsReponse
}
