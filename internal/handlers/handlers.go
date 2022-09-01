/**
*	NAME: Aaron Meek
*	DATE: 2022 - 08 - 24
*
*	This contains the page handlers
 */
package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/urhumantoast/bookings/internal/config"
	"github.com/urhumantoast/bookings/internal/forms"
	"github.com/urhumantoast/bookings/internal/models"
	"github.com/urhumantoast/bookings/internal/render"
)

type Repository struct {
	App *config.AppConfig
}

// Repo - The repository used by the handlers
var Repo *Repository

// NewRepo - Creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers - Sets the repository for handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// About - The about page renderer
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIp

	// Send the data to the template page (as a test), and render the page
	render.RenderTemplate(w, r, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

// SmallRooms - The small rooms page renderer
func (m *Repository) SmallRooms(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "small-rooms.page.html", &models.TemplateData{})
}

// MediumRooms - The medium rooms page renderer
func (m *Repository) MiddleRooms(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "middle-rooms.page.html", &models.TemplateData{})
}

// LargeRooms - The large rooms page renderer
func (m *Repository) LargeRooms(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "large-rooms.page.html", &models.TemplateData{})
}

// Support - The support page renderer
func (m *Repository) Support(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "support.page.html", &models.TemplateData{})
}

// Reservations - The Reservations page renderer
func (m *Repository) Reservations(w http.ResponseWriter, r *http.Request) {

	var emptyReservation models.Reservation
	data := make(map[string]interface{})
	data["reservation"] = emptyReservation

	render.RenderTemplate(w, r, "reservations.page.html", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

// PostReservation - Handles posting of a reservation form
func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	// Get the user data entered in the form
	reservation := models.Reservation{
		FirstName:    r.Form.Get("first-name"),
		LastName:     r.Form.Get("last-name"),
		EmailAddress: r.Form.Get("email-address"),
		PhoneNumber:  r.Form.Get("phone-number"),
	}

	form := forms.New(r.PostForm)

	// Run validation check on the data
	form.Required("first-name", "last-name", "email-address", "phone-number")
	form.MinLength("first-name", 3, r)
	form.MinLength("last-name", 3, r)
	form.Match("phone-number", "[0-9]{3}[-][0-9]{3}[-][0-9]{4}", "555-555-5555", r)
	form.IsEmail("email-address")

	// If the form isnt valid, notify the user
	if !form.Valid() {
		// Record the valid data
		data := make(map[string]interface{})
		data["reservation"] = reservation

		// Render the page
		render.RenderTemplate(w, r, "reservations.page.html", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

	m.App.Session.Put(r.Context(), "reservation", reservation)

	// Direct user to reservation summary (also prevents duplicate submition)
	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)
}

// Availability - The Availability page renderer
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-availability.page.html", &models.TemplateData{})
}

// PostAvailability - Handles post for Availability page
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")
	w.Write([]byte(fmt.Sprintf("start date is %s and end date is %s", start, end)))
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// AvailabilityJSON - Handles request for Availability page and sends JSON reponse
func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      true,
		Message: "TESTING JSON",
	}

	out, err := json.MarshalIndent(resp, "", "     ")
	if err != nil {
		log.Println(err)
	}

	log.Println(string(out))
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func (m *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {
	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		m.App.Session.Put(r.Context(), "error", "Reservation summary not available")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	m.App.Session.Remove(r.Context(), "reservation")
	data := make(map[string]interface{})
	data["reservation"] = reservation

	render.RenderTemplate(w, r, "reservations-summary.page.html", &models.TemplateData{
		Data: data,
	})
}

// Home - The home page renderer
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "home.page.html", &models.TemplateData{})
}
