package main

import (
	"Lab7_RestAPI/packages/handlers"
	"fmt"
	"net/http"
)

const (
	QUADRATIC_EQUATION_ROUTE = "/quadratic-equation/"
	SLICE_ROUTE              = "/slice/"
	SERVER_URL               = "http://localhost"
	SEVER_PORT               = ":8080"
)

func main() {
	run()
}

func run() {

	mux := http.NewServeMux()

	mux.HandleFunc(QUADRATIC_EQUATION_ROUTE, handlers.QuadraticEquation)
	mux.HandleFunc(SLICE_ROUTE, handlers.Slice)
	mux.HandleFunc("/", handlers.Index)

	fmt.Println("Server is launching on " + SERVER_URL + SEVER_PORT)
	err := http.ListenAndServe(SEVER_PORT, mux)
	if err != nil {
		fmt.Println("An error occurred while starting the server", "\n", err.Error())
		return
	}
}
