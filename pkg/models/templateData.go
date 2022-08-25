/**
*	NAME: Aaron Meek
*	DATE: 2022 - 08 - 24
*
*	This contains the template data/page struct
 */
package models

// TemplateData - Holds data sent from handlers to template
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}
