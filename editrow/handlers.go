package editrow

import (
	"net/http"
	"strconv"
	"strings"
)

type user struct {
	name, email string
}

var users = []user{
	{"Joe Smith", "joe@smith.org"},
	{"Angie MacDowell", "angie@macdowell.org"},
	{"Fuqua Tarkenton", "fuqua@tarkenton.org"},
	{"Kim Yee", "kim@yee.org"},
}

func Handlers(prefix string, mux *http.ServeMux) {
	mux.HandleFunc(prefix+"/", index)
	mux.HandleFunc(prefix+"/edit/", editUser)
	mux.HandleFunc(prefix+"/contact/", contacts)
}

func index(w http.ResponseWriter, r *http.Request) {
	// Load users
	Index(users).Render(r.Context(), w)
}

func editUser(w http.ResponseWriter, r *http.Request) {
	segments := strings.Split(r.URL.Path, "/")
	userID, _ := strconv.Atoi(segments[len(segments)-1])
	form(userID, users[userID]).Render(r.Context(), w)
}

func contacts(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		getRow(w, r)
		return
	}
	if r.Method == http.MethodPut {
		putUser(w, r)
		return
	}
	w.WriteHeader(404)
}

func getRow(w http.ResponseWriter, r *http.Request) {
	segments := strings.Split(r.URL.Path, "/")
	userID, _ := strconv.Atoi(segments[len(segments)-1])
	row(userID, users[userID]).Render(r.Context(), w)
}

func putUser(w http.ResponseWriter, r *http.Request) {
	segments := strings.Split(r.URL.Path, "/")
	userID, _ := strconv.Atoi(segments[len(segments)-1])
	r.ParseForm()
	name, email := r.FormValue("name"), r.FormValue("email")
	users[userID] = user{name, email}
	row(userID, users[userID]).Render(r.Context(), w)
}
