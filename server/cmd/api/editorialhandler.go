package main

import (
	"biblioteca/internal/data"
	"net/http"
)

func (app *application) createEditorialHandler(w http.ResponseWriter, r *http.Request) {

	var input  data.Editorial

	err := app.readJSON(w , r , &input)
	if err != nil {
		app.badRequestResponse(w , r , err)
		return
	}

	editorial := &data.Editorial{
		Nombre: input.Nombre,
		Direccion:  input.Direccion,
		PaginaWeb:  input.PaginaWeb,
	}

	err = app.models.Editorial.CreateEditorial( *editorial)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusAccepted, envelope{"editorial": editorial}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}


func (app *application) getEditorialsHandler(w http.ResponseWriter, r *http.Request) {
	

	editorial, err := app.models.Editorial.GetEditorials()

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	
	err = app.writeJSON(w, http.StatusAccepted, envelope{"editorials": editorial}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}