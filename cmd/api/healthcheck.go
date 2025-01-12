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
