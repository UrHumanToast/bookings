package handlers

import (
	"net/http"

	"github.com/urhumantoast/bookings/pkg/config"
	"github.com/urhumantoast/bookings/pkg/models"
	"github.com/urhumantoast/bookings/pkg/render"
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

// Home - The home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

// About - The about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIp

	// Send the data to the template page, and render the page
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
