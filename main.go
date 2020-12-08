package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nicobianchetti/Middleware-Go/controller"
	"github.com/nicobianchetti/Middleware-Go/functions"
	"github.com/nicobianchetti/Middleware-Go/middleware"
)

func execute(name string, f functions.MyFunction) {
	f(name)
}

func main() {
	/*
		name := "Proyecto usando middlewarewe"
		// execute(name, functions.Saludar)
		// execute(name, functions.Despedirse)

		execute(name, functions.MiddlewareLog(functions.Saludar))
		execute(name, functions.MiddlewareLog(functions.Despedirse))
	*/

	router := mux.NewRouter()

	// router.HandleFunc("/saludar", controller.GetSaludo).Methods(http.MethodGet)

	//envuelvo el handler con el middleware
	router.HandleFunc("/saludar", middleware.Log(middleware.Authentication(controller.GetSaludo))).Methods(http.MethodGet)

	http.ListenAndServe(":8080", router)

}
