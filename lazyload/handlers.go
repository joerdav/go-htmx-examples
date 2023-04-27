package lazyload

import (
	_ "embed"
	"net/http"
	"time"
)

var (
	//go:embed tokyo.png
	graphBytes []byte
	//go:embed bars.svg
	barsBytes []byte
)

func Handlers(prefix string, mux *http.ServeMux) {
	mux.HandleFunc(prefix+"/", index)
	mux.HandleFunc(prefix+"/graph", graph)
	mux.HandleFunc(prefix+"/bars.svg", bars)
	mux.HandleFunc(prefix+"/tokyo.png", tokyo)
}

func index(w http.ResponseWriter, r *http.Request) {
	Index().Render(r.Context(), w)
}

func graph(w http.ResponseWriter, r *http.Request) {
	time.Sleep(2 * time.Second)
	graphImage().Render(r.Context(), w)
}

func bars(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	w.Write(barsBytes)
}

func tokyo(w http.ResponseWriter, r *http.Request) {
	w.Write(graphBytes)
}
