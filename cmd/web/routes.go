package main

import (
	"net/http"

	"github.com/urhumantoast/bookings/pkg/config"
	"github.com/urhumantoast/bookings/pkg/handlers"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/small-rooms", handlers.Repo.SmallRooms)
	mux.Get("/middle-rooms", handlers.Repo.MiddleRooms)
	mux.Get("/large-rooms", handlers.Repo.LargeRooms)
	mux.Get("/support", handlers.Repo.Support)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
