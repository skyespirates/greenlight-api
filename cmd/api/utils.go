package main

import "greenlight.skyespirates.net/internal/data"

func filter(slice *[]data.Movie, condition func(data.Movie) bool) {
	writeIndex := 0
	source := *slice
	for _, value := range source {
		if condition(value) {
			source[writeIndex] = value
			writeIndex++
		}
	}
	*slice = source[:writeIndex]
}

func findMovieById(movies *[]data.Movie, id int64) (data.Movie, bool) {
	for _, movie := range *movies {
		if movie.ID == id {
			return movie, false
		}
	}
	return data.Movie{}, true
}
