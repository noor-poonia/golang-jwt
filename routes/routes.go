package routes

import (
	"github.com/gorilla/mux"
	"jwt/handlers"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	
	router.HandleFunc("/login", handlers.Login).Methods("GET")
	router.HandleFunc("/home", handlers.Home).Methods("POST")
	router.HandleFunc("/refresh", handlers.Refresh).Methods("PUT")
	
	return router
}