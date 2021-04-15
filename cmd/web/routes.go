package main

import (
	config2 "github.com/firmanali/book-reservation/internal/config"
	handlers2 "github.com/firmanali/book-reservation/internal/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config2.AppConfig) http.Handler {

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers2.Repo.Home)
	mux.Get("/about", handlers2.Repo.About)
	mux.Get("/generals-quarter", handlers2.Repo.Generals)
	mux.Get("/majors-suite", handlers2.Repo.Majors)
	mux.Get("/make-reservation", handlers2.Repo.Reservation)
	mux.Get("/search-availability", handlers2.Repo.Availability)
	mux.Post("/search-availability", handlers2.Repo.PostAvailability)
	mux.Post("/search-availability-json", handlers2.Repo.AvailabilityJson)
	mux.Get("/contact", handlers2.Repo.Contact)

	fileserver := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileserver))

	return mux
}
