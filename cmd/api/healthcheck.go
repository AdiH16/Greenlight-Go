package main

import (
	"fmt"
	"net/http"
)

//Handler that writes a plain text response with information about
//application status, operating environment and version

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "environment: %s\n", app.config.env)
	fmt.Fprintf(w, "version: %s\n", version)
}
