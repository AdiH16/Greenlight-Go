package main

import (
	"fmt"
	"net/http"
)

//Handler that writes a plain text response with information about
//application status, operating environment and version

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	// fixed-format JSON response from a string
	js := `{"status": "available", "environment": %q, "version": %q}`
	js = fmt.Sprintf(js, app.config.env, version)

	w.Header().Set("Content-type", "application/json")

	w.Write([]byte(js))
}
