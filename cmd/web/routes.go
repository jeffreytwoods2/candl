package main

import (
	"net/http"

	"candl.jwoods.dev/ui"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /static/", http.FileServerFS(ui.Files))

	mux.HandleFunc("GET /", app.homeHandler)

	return mux
}
