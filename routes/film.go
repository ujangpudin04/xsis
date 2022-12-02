package routes

import (
	"xsis/handlers"
	"xsis/pkg/mysql"
	"xsis/repositories"

	"github.com/gorilla/mux"
)

func FilmRoutes(r *mux.Router) {
	FilmRepository := repositories.RepositoryFilm(mysql.DB)
	h := handlers.HandlerFilm(FilmRepository)

	r.HandleFunc("/film", h.FindFilm).Methods("GET")
	r.HandleFunc("/film/{id}", h.GetFilm).Methods("GET")
	r.HandleFunc("/film", h.CreateFilm).Methods("POST")
	// r.HandleFunc("/film", middleware.Auth(middleware.UploadFile(h.CreateFilm, "thumbnailfilm"))).Methods("POST")
	r.HandleFunc("/film/{id}", h.UpdateFilm).Methods("PATCH")
	r.HandleFunc("/film/{id}", h.DeleteFilm).Methods("DELETE")
}
