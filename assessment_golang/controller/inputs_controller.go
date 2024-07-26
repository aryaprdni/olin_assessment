package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type InputsController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
