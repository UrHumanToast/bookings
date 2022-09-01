package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/asaskevich/govalidator"
)

// Creates a custom form struct, embeds a url.Values object
type Form struct {
	url.Values
	Errors errors
}

// New - Initializes a form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Required - Variadic: to check for data in required fields
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "Please fill")
		}
	}
}

// Has - Checks if form field is in post and not empty
// func (f *Form) Has(field string, r *http.Request) bool {
// 	if len(r.Form.Get(field)) == 0 {
// 		return false
// 	}

// 	return true
// }

func (f *Form) MinLength(field string, length int, r *http.Request) bool {
	x := r.Form.Get(field)
	if len(x) < length {
		f.Errors.Add(field, fmt.Sprintf("minimum %d letters", length))
		return false
	}

	return true
}

// Match - Uses regular expressions to match a string to a pattern
func (f *Form) Match(field string, pattern string, format string, r *http.Request) bool {
	x := r.Form.Get(field)
	expr, _ := regexp.Compile(pattern)
	if !(expr.MatchString(x)) {
		f.Errors.Add(field, fmt.Sprintf("Format: %s", format))
		return false
	}

	return true
}

// IsEmail - Checks for a valid email address
func (f *Form) IsEmail(field string) bool {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "Invalid Email")
		return false
	}

	return true
}

// Valid - Returns true if there are no errors, otherwise false
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
