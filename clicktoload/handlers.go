package clicktoload

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/rs/xid"
)

type user struct {
	name, email, id string
}

func Handlers(prefix string, mux *http.ServeMux) {
	mux.HandleFunc(prefix+"/", index)
	mux.HandleFunc(prefix+"/contacts/", getPage)
}

func index(w http.ResponseWriter, r *http.Request) {
	// Load users
	Index(getUsers(1)).Render(r.Context(), w)
}

func getPage(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	tbody(getUsers(page), page).Render(r.Context(), w)
}

func getUsers(page int) []user {
	var users []user
	for i := 0; i < 10; i++ {
		users = append(users, user{
			"Agent Smith",
			fmt.Sprintf("void%d@null.org", (page*10)+i),
			xid.NewWithTime(time.Now()).String(),
		})
	}
	return users
}
