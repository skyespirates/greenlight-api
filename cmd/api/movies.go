package main

import (
	"fmt"
	"net/http"
	"time"

	"greenlight.skyespirates.net/internal/data"
)

var movies = []data.Movie{
	{
		ID:        1,
		CreatedAt: time.Now(),
		Title:     "Casablanca",
		Runtime:   102,
		Genres:    []string{"drama", "romance", "war"},
		Version:   1,
	},
	{
		ID:        2,
		CreatedAt: time.Now(),
		Title:     "Cool Hand Luke",
		Runtime:   126,
		Genres:    []string{"crime", "drama"},
		Version:   1,
	},
	{
		ID:        3,
		CreatedAt: time.Now(),
		Title:     "Bullitt",
		Runtime:   114,
		Genres:    []string{"action", "crime", "thriller"},
		Version:   1,
	},
}

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title string `json:"title"`
		Year int32 `json:"year"`
		Runtime data.Runtime `json:"runtime"`
		Genres []string `json:"genres"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	fmt.Fprintf(w, "%+v\n", input)
	// app.writeJSON(w, http.StatusOK, envelope{"movie": input}, nil)
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 0 {
		app.notFoundResponse(w, r)
		return
	}
	movie, isNotFound := findMovieById(&movies, id)
	if isNotFound == true {
		app.notFoundResponse(w, r)
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getAllMoviesHandler(w http.ResponseWriter, r *http.Request) {
	err := app.writeJSON(w, http.StatusOK, envelope{"movies": movies}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 0 || id > int64(len(movies)-1) {
		app.notFoundResponse(w, r)
		return
	}
	filter(&movies, func(movie data.Movie) bool {
		return movie.ID != id
	})
	err = app.writeJSON(w, http.StatusOK, envelope{"message": "movie deleted successfully"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
