package main

import (
	"fmt"
	"net/http"
	"time"

	"greenlight.adih.net/internal/data"
)

// createMovieHandler for "POST /v1/movies" enpoint
func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	// anonymous struct to hold the information we expect in the HTTP request body
	var input struct {
		Title   string       `json:"title"`
		Year    int32        `json:"year"`
		Runtime data.Runtime `json:"runtime"`
		Genres  []string     `json:"genres"`
	}

	// use the readJSON() to decode the request body into input struct
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
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
