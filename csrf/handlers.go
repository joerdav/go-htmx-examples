package csrf

import (
	"log"
	"net/http"

	"github.com/justinas/nosurf"
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
	mux.HandleFunc(prefix, index)
	mux.Handle(prefix+"/contact/1", nosurf.New(http.HandlerFunc(putUser)))
	mux.Handle(prefix+"/contact/1/edit", nosurf.New(http.HandlerFunc(editForm)))
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
	Form(user, nosurf.Token(r)).Render(r.Context(), w)
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
