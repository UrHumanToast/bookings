/**
*	NAME: Aaron Meek
*	DATE: 2022 - 08 - 24
*
*	This contains the routing
 */
package main

import (
	"net/http"

	"github.com/urhumantoast/bookings/internal/config"
	"github.com/urhumantoast/bookings/internal/handlers"

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
	mux.Get("/support", handlers.Repo.Support)

	mux.Get("/small-rooms", handlers.Repo.SmallRooms)
	mux.Get("/middle-rooms", handlers.Repo.MiddleRooms)
	mux.Get("/large-rooms", handlers.Repo.LargeRooms)

	mux.Get("/book-reservation", handlers.Repo.Reservations)
	mux.Post("/book-reservation", handlers.Repo.PostReservation)
	mux.Get("/reservation-summary", handlers.Repo.ReservationSummary)

	mux.Get("/search-availability", handlers.Repo.Availability)
	mux.Post("/search-availability", handlers.Repo.PostAvailability)
	mux.Post("/search-availability-json", handlers.Repo.AvailabilityJSON)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
