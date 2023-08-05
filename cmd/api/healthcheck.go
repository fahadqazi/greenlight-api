package main

import (
	"fmt"
	"net/http"
)

func (a *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "environtment: %s\n", a.config.env)
	fmt.Fprintf(w, "version: %s\n", version)
}

func (a *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create movie")
}

func (a *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "show movie")
}
