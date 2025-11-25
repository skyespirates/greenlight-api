package main

import (
	"errors"
	"net/http"
	"time"

	"greenlight.skyespirates.net/internal/data"
	"greenlight.skyespirates.net/internal/validator"
)

func (app *application) registerUserHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	user := &data.User{
		Name:      input.Name,
		Email:     input.Email,
		Activated: false,
	}
	err = user.Password.Set(input.Password)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()
	if data.ValidateUser(v, user); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	err = app.models.Users.Insert(user)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrDuplicateEmail):
			v.AddError("email", "a user with this email address already exists")
			app.failedValidationResponse(w, r, v.Errors)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	token, err := app.models.Tokens.New(user.ID, 3*24*time.Hour, data.ScopeActivation)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	app.background(func() {
		dt := map[string]interface{}{
			"activationToken": token.Plaintext,
			"user":            user,
		}
		err = app.mailer.Send(user.Email, "user_welcome.tmpl", dt)
		if err != nil {
			app.logger.PrintError(err, nil)
		}
	})

	app.successResponse(w, http.StatusAccepted, "user registered successfully", envelope{"user": user})
	// err = app.writeJSON(w, http.StatusAccepted, envelope{"user": user}, nil)
	// if err != nil {
	// 	app.serverErrorResponse(w, r, err)
	// }
}

func (app *application) activateUserHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		TokenPlainText string `json:"token"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
	}

	v := validator.New()
	if data.ValidateTokenPlaintext(v, input.TokenPlainText); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	user, err := app.models.Users.GetForToken(data.ScopeActivation, input.TokenPlainText)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			v.AddError("token", "invalid or expired activation token")
			app.failedValidationResponse(w, r, v.Errors)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	user.Activated = true
	err = app.models.Users.Update(user)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrEditConflict):
			app.editConflictResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.models.Tokens.DeleteAllForUser(data.ScopeActivation, user.ID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.successResponse(w, http.StatusOK, "user activated successfully", envelope{"user": user})
	// err = app.writeJSON(w, http.StatusOK, envelope{"user": user}, nil)
	// if err != nil {
	// 	app.serverErrorResponse(w, r, err)
	// }
}

func (app *application) changePasswordHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		CurrentPassword    string `json:"current_password"`
		NewPassword        string `json:"new_password"`
		ConfirmNewPassword string `json:"confirm_new_password"`
	}

	user := app.contextGetUser(r)

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	// validate password
	// make sure current password corrent
	// ensure newPassword and CorfirmNewPassword same
	match, err := user.Password.Matches(input.CurrentPassword)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	if !match {
		app.invalidCredentialsResponse(w, r)
		return
	}

	v := validator.New()
	data.ValidatePasswordPlaintext(v, input.NewPassword)

	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	if input.NewPassword != input.ConfirmNewPassword {
		app.badRequestResponse(w, r, errors.New("new passwords are not equal"))
		return
	}

	err = app.models.Users.ChangePassword(user, input.NewPassword)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.successResponse(w, http.StatusAccepted, "password changed successfully", envelope{"user": user})
	// app.writeJSON(w, http.StatusAccepted, envelope{"user": user}, nil)
}
