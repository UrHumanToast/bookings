/**
*	NAME: Aaron Meek
*	DATE: 2022 - 08 - 24
*
*	This contains the application configuration struct
 */
package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
}
