// package main

// import (
// 	"aryaprdni/assessment_golang/app"
// 	"aryaprdni/assessment_golang/controller"
// 	"aryaprdni/assessment_golang/helper"
// 	"aryaprdni/assessment_golang/repository"
// 	"aryaprdni/assessment_golang/service"
// 	"net/http"

// 	_ "github.com/lib/pq"

// 	"github.com/go-playground/validator/v10"
// 	"github.com/julienschmidt/httprouter"
// )

// func main() {

// 	db := app.NewDB()
// 	validate := validator.New()
// 	inputsRepository := repository.NewInputsRepository()
// 	inputsService := service.NewInputsService(inputsRepository, db, validate)
// 	inputController := controller.NewInputsController(inputsService)

// 	router := httprouter.New()

// 	router.POST("/api/twosum", inputController.Create)

// 	server := http.Server{
// 		Addr:    "localhost:8080",
// 		Handler: router,
// 	}

// 	err := server.ListenAndServe()
// 	helper.PanicIfError(err)
// }
