package main

import (
	"net/http"
	"time"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	info := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}

	// demonstrate graceful shutdown
	time.Sleep(4 * time.Second)

	data := map[string]interface{}{
		"info": info,
	}

	app.successResponse(w, http.StatusOK, "health check information", data)
}
