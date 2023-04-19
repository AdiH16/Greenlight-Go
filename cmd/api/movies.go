package main

import (
	"fmt"
	"net/http"
	"time"

	"greenlight.adih.net/internal/data"
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
		app.notFoundResponse(w, r)
		return
	}

	// instance of the movie struct with some dummy data
	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Casablanca",
		Runtime:   102,
		Genres:    []string{"drama", "romance", "war"},
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
