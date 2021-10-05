package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"redoc-sample/api"
)

func main() {
	ctrl := api.NewItemsController()

	sm := mux.NewRouter()

	getR := sm.Methods(http.MethodGet).Subrouter()
	getR.HandleFunc("/items", ctrl.GetAll)
	getR.HandleFunc("/items/{id:[0-9]+}", ctrl.GetOne)

	postR := sm.Methods(http.MethodPost).Subrouter()
	postR.HandleFunc("/items", ctrl.Create)

	putR := sm.Methods(http.MethodPut).Subrouter()
	putR.HandleFunc("/items/{id:[0-9]+}", ctrl.Update)

	deleteR := sm.Methods(http.MethodDelete).Subrouter()
	deleteR.HandleFunc("/items/{id:[0-9]+}", ctrl.Delete)

	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	getR.Handle("/docs", sh)
	getR.Handle("/swagger.yaml", http.FileServer(http.Dir("./docs")))

	s := http.Server{
		Addr:         ":2805",
		Handler:      sm,
	}

	fmt.Println("Run 'go get -u github.com/go-swagger/go-swagger/cmd/swagger'")
	fmt.Println("Run 'swagger generate spec -o ./docs/swagger.yaml --scan-models'")
	fmt.Println("Go 'http://localhost:2805/docs'")

	err := s.ListenAndServe()
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}
}
