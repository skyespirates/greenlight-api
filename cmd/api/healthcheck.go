package main

import (
	"net/http"
	"time"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}

	// demonstrate graceful shutdown
	time.Sleep(4 * time.Second)

	err := app.writeJSON(w, http.StatusOK, envelope{"healthcheck": data}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) indexHandler(w http.ResponseWriter, r *http.Request) {
	// data := map[string]interface{}{
	// 	"about": "inspired by TMDB (The Movie Database)",
	// 	"usage": "manage movies information",
	// }

	dt := struct {
		About string `json:"about"`
		Usage string `json:"usage"`
	}{
		About: "inspired by TMDB (The Movie Database)",
		Usage: "where you manage movies information",
	}

	app.successResponse(w, http.StatusOK, "manage movies information", dt)
}
