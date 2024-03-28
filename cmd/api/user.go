package main

import (
	"errors"
	"net/http"

	"github.com/FMinister/co2monitor-api/internal/data"
	"github.com/FMinister/co2monitor-api/internal/validator"
)

func (app *application) getUserHandler(w http.ResponseWriter, r *http.Request) {
	qs := r.URL.Query()
	v := validator.New()
	name := app.readString(qs, "name", "")
	id := app.readInt(qs, "id", 0, v)

	if name == "" && id == 0 {
		app.badRequestResponse(w, r, errors.New("name or id parameter is missing"))
		return
	}
	if name != "" && id != 0 {
		app.badRequestResponse(w, r, errors.New("only one of name or id parameter is allowed"))
		return
	}

	var user data.User
	var err error

	if id != 0 {
		user, err = app.models.User.GetByID(id)
	}
	if name != "" {
		user, err = app.models.User.GetByName(name)
	}

	switch err {
	case nil:
	case data.ErrRecordNotFound:
		app.notFoundResponse(w, r)
		return
	default:
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"user": user}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) createUserHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name   string `json:"name"`
		Token  string `json:"token"`
		Active bool   `json:"active"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	user := &data.User{
		Name:   input.Name,
		Token:  input.Token,
		Active: input.Active,
	}

	v := validator.New()
	if data.ValidateUser(v, user); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.User.Insert(user)
	if err != nil {
		switch err {
		case data.ErrDuplicateName:
			v.AddError("name", "a user with this name already exists")
			app.failedValidationResponse(w, r, v.Errors)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{"user": user}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) updateUserHandler(w http.ResponseWriter, r *http.Request) {
	qs := r.URL.Query()
	v := validator.New()
	name := app.readString(qs, "name", "")
	id := app.readInt(qs, "id", 0, v)

	if name == "" && id == 0 {
		app.badRequestResponse(w, r, errors.New("name or id parameter is missing"))
		return
	}
	if name != "" && id != 0 {
		app.badRequestResponse(w, r, errors.New("only one of name or id parameter is allowed"))
		return
	}

	var user data.User
	var err error

	if id != 0 {
		user, err = app.models.User.GetByID(id)
	}
	if name != "" {
		user, err = app.models.User.GetByName(name)
	}

	switch err {
	case nil:
	case data.ErrRecordNotFound:
		app.notFoundResponse(w, r)
		return
	default:
		app.serverErrorResponse(w, r, err)
		return
	}

	var input struct {
		Name   *string `json:"name"`
		Token  *string `json:"token"`
		Active *bool   `json:"active"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if input.Name != nil {
		user.Name = *input.Name
	}
	if input.Token != nil {
		user.Token = *input.Token
	}
	if input.Active != nil {
		user.Active = *input.Active
	}

	if data.ValidateUser(v, &user); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.User.Update(&user)
	if err != nil {
		switch err {
		case data.ErrDuplicateName:
			v.AddError("name", "a user with this name already exists")
			app.failedValidationResponse(w, r, v.Errors)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"user": user}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.models.User.Delete(id)
	if err != nil {
		switch err {
		case data.ErrRecordNotFound:
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusNoContent, nil, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
