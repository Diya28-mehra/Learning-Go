package router

import (
	"github.com/Diya28-mehra/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router{
	router := mux.NewRouter()
	router.HandleFunc("/api/movies",controllers.GetMyAllMovies).Methods("GET")
	router.HandleFunc("/api/movie",controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}",controllers.MarkedAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}",controllers.DeleteAMovie).Methods("DELETE")
	router.HandleFunc("/api/deleteall",controllers.DeleteAllMovies).Methods("DELETE")
	return router
}