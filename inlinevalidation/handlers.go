package inlinevalidation

import (
	"errors"
	"net/http"

	"github.com/a-h/pathvars"
)

var validateMatcher = pathvars.NewExtractor("/inline-validation/validate/{name}")

func Handlers(prefix string, mux *http.ServeMux) {
	mux.HandleFunc(prefix+"/", index)
	mux.HandleFunc(prefix+"/validate/", validate)
}

func index(w http.ResponseWriter, r *http.Request) {
	Index().Render(r.Context(), w)
}

func validate(w http.ResponseWriter, r *http.Request) {
	vals, ok := validateMatcher.Extract(r.URL)
	if !ok {
		w.WriteHeader(500)
		return
	}
	name, ok := vals["name"]
	if !ok {
		w.WriteHeader(500)
		return
	}
	f, ok := fields[name]
	if !ok {
		w.WriteHeader(500)
		return
	}
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(500)
		return
	}
	val := r.FormValue(name)
	inp(f, name, val, f.validator(val)).Render(r.Context(), w)
}

type field struct {
	text      string
	validator func(string) error
}

var fields = map[string]field{
	"email": {
		text: "Email",
		validator: func(value string) error {
			if value != "test@test.com" {
				return errors.New("Only test@test.com is valid.")
			}
			return nil
		},
	},
	"firstName": {
		text: "First Name",
		validator: func(value string) error {
			if value == "" {
				return errors.New("Required")
			}
			return nil
		},
	},
	"lastName": {
		text: "Last Name",
		validator: func(value string) error {
			if value == "" {
				return errors.New("Required")
			}
			return nil
		},
	},
}
