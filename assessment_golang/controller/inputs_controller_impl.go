package controller

import (
	"aryaprdni/assessment_golang/helper"
	"aryaprdni/assessment_golang/model/web"
	"aryaprdni/assessment_golang/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type InputsControllerImpl struct {
	InputsService service.InputsService
}

func NewInputsController(inputsService service.InputsService) InputsController {
	return &InputsControllerImpl{
		InputsService: inputsService,
	}
}

func (controller *InputsControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	inputsCreateRequest := web.InputsCreateRequest{}
	helper.ReadFromRequestBody(request, &inputsCreateRequest)

	inputsReponse := controller.InputsService.Create(request.Context(), inputsCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   inputsReponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
