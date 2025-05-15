package main

import (
	"biblioteca/internal/data"
	"net/http"
)

// getAutoresHandler godoc
// @Sumary  Authors  
// @Tags    Authors 
// @Accept  json 
// @Produce json 
// @Router  /api/autores [get] 
func (app *application) getAutoresHandler(w http.ResponseWriter, r *http.Request) {

	authors , err  := app.models.Autor.GetAutores()
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusAccepted, envelope{"user": authors}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) createAutorHander(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name        string `json:"name"`
		Nationality string `json:"nationality"`
	}

	err :=  app.readJSON(w , r , &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	author := &data.Author{
		Nombre: input.Name,
		Nacionalidad: input.Nationality,
	}

	err = app.models.Autor.CreateAuthor(author)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusAccepted, envelope{"user": author}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
