package clicktoedit

import (
	"log"
	"net/http"
)

type user struct {
	firstName, lastName, email string
}

var demoUser user = user{
	firstName: "Joe",
	lastName:  "Blow",
	email:     "joe@blow.com",
}

func Handlers(prefix string, mux *http.ServeMux) {
	mux.HandleFunc(prefix+"/", index)
	mux.HandleFunc(prefix+"/contact/1", putUser)
	mux.HandleFunc(prefix+"/contact/1/edit", editForm)
}

func index(w http.ResponseWriter, r *http.Request) {
	// Load user
	user := demoUser
	if r.Header.Get("HX-Request") == "true" {
		Display(user).Render(r.Context(), w)
		return
	}
	Index(user).Render(r.Context(), w)
}

func editForm(w http.ResponseWriter, r *http.Request) {
	// Load user
	user := demoUser
	Form(user).Render(r.Context(), w)
}

func putUser(w http.ResponseWriter, r *http.Request) {
	// Load user
	if err := r.ParseForm(); err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	demoUser = user{
		firstName: r.FormValue("firstName"),
		lastName:  r.FormValue("lastName"),
		email:     r.FormValue("email"),
	}
	Display(demoUser).Render(r.Context(), w)
}
