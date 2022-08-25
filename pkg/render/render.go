/**
*	NAME: Aaron Meek
*	DATE: 2022 - 08 - 24
*
*	This contains the page rendering logic
 */
package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/urhumantoast/bookings/pkg/config"
	"github.com/urhumantoast/bookings/pkg/models"
)

var app *config.AppConfig

// NewTemplates - Sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// RenderTemplate - Renders template using HTML template
func RenderTemplate(w http.ResponseWriter, tmpl_name string, td *models.TemplateData) {
	var tc map[string]*template.Template

	// Get the template cache from the app config
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		var err error
		tc, err = CreateTemplateCache()
		if err != nil {
			log.Println(err)
		}
	}

	// Get requested template
	t, ok := tc[tmpl_name]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	// Render the page
	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	err := t.Execute(buf, td)
	if err != nil {
		log.Println(err)
	}

	// Render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// Get all of the files named *.page.html from the ./templates folder
	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	// Range through all files named *.page.html
	for _, page := range pages {
		// Strip off the path to the file
		name := filepath.Base(page)

		// Give the template a name and parse its respective page
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		// Check for a layout file
		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		// Parse any layouts found
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}

		}

		// Add to the template cache
		myCache[name] = ts
	}

	return myCache, nil
}
