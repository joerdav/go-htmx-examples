package main

import (
	"examples/bulkupdate"
	"examples/clicktoedit"
	"examples/clicktoload"
	"log"
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	r := http.NewServeMux()
	r.Handle("/", templ.Handler(Home(examples)))
	for _, e := range examples {
		log.Printf("Serving %q on /%s", e.Name, e.Slug)
		e.Handlers("/"+e.Slug, r)
	}
	log.Println("Listening on localhost:2468")
	return http.ListenAndServe("localhost:2468", r)
}

type Example struct {
	Name, Desc, Slug string
	Handlers         func(string, *http.ServeMux)
}

var examples = []Example{
	{
		Name:     "Click To Edit",
		Desc:     "Demonstrates inline editing of a data object",
		Slug:     "click-to-edit",
		Handlers: clicktoedit.Handlers,
	},
	{
		Name:     "Bulk Update",
		Desc:     "Demonstrates bulk updating of multiple rows of data",
		Slug:     "bulk-update",
		Handlers: bulkupdate.Handlers,
	},
	{
		Name:     "Click to Load",
		Desc:     "Demonstrates clicking to load more rows in a table",
		Slug:     "click-to-load",
		Handlers: clicktoload.Handlers,
	},
}
