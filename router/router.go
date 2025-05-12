package router

import (
	"module/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/movies", controller.Cgetall).Methods("GET")
	r.HandleFunc("/api/movie/{id}", controller.Cgetspecificw).Methods("GET")
	r.HandleFunc("/api/movie", controller.Cadd).Methods("POST")
	r.HandleFunc("/api/movie/{id}",controller.Cupdate ).Methods("PUT")
   return r;
}
