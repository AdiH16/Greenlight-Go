package main

import (
	"fmt"
	"net/http"
)

// createMovieHandler for "POST /v1/movies" enpoint
func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "create a new movie")
}

//showMovieHandler for "GET /v1/movies/:id" endpoint
//Retrieve the "id" parameter from the URL

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {

	//get value of the "id" with ByName() and convert it
	//if the parameter cannot be converted it is invalid
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	//interpolate movie ID in a placeholder response
	fmt.Fprintf(w, "show the details of movie %d\n", id)
}
