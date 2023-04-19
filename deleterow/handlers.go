package deleterow

import (
	"fmt"
	"net/http"
	"strings"
)

type user struct {
	name, email, id string
}

func Handlers(prefix string, mux *http.ServeMux) {
	mux.HandleFunc(prefix+"/", index)
	mux.HandleFunc(prefix+"/contact/", deleteUser)
}

func index(w http.ResponseWriter, r *http.Request) {
	// Load users
	Index().Render(r.Context(), w)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	segments := strings.Split(r.URL.Path, "/")
	userID := segments[len(segments)-1]
	fmt.Println("Delete user: ", userID)
}
