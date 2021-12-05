package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yusuftatli/hepsiburada/common"
	"github.com/yusuftatli/hepsiburada/handlers"
)


func InitializeRoutes(env *common.Environment) {

	api := mux.NewRouter().PathPrefix("/").Subrouter()
	api.HandleFunc("/currency/{currencyCode}", handlers.GetCurrency).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8081",api))
}