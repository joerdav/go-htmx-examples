package bulkupdate

import (
	"log"
	"net/http"
	"strconv"
)

type user struct {
	name, email string
	active      bool
}

var inMemDB []user = []user{
	{"Joe Smith", "joe@smith.org", true},
	{"Angie MacDowell", "angie@macdowell.org", true},
	{"Fuqua Tarketon", "fuqua@tarketon.org", true},
	{"Kim Yee", "kim@yee.org", false},
}

func Handlers(prefix string, mux *http.ServeMux) {
	mux.HandleFunc(prefix+"/", index)
	mux.HandleFunc(prefix+"/activate", putActivate)
	mux.HandleFunc(prefix+"/deactivate", putDeactivate)
}

func index(w http.ResponseWriter, r *http.Request) {
	// Load users
	Index(inMemDB).Render(r.Context(), w)
}

func putActivate(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	ids := map[int]bool{}
	for _, i := range r.Form["ids"] {
		id, _ := strconv.Atoi(i)
		user := inMemDB[id]
		user.active = true
		inMemDB[id] = user
		ids[id] = true
	}
	tbody(inMemDB, ids).Render(r.Context(), w)
}

func putDeactivate(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	ids := map[int]bool{}
	for _, i := range r.Form["ids"] {
		id, _ := strconv.Atoi(i)
		user := inMemDB[id]
		user.active = false
		inMemDB[id] = user
		ids[id] = true
	}
	tbody(inMemDB, ids).Render(r.Context(), w)
}
