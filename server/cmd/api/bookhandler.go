package main

import (
	"biblioteca/internal/data"
	"encoding/json"
	"net/http"
)

func (app *application) createBookHandler(w http.ResponseWriter, r *http.Request) {

	var input data.Book

	err := app.readJSON(w, r, &input)

	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	book := &data.Book{
		Titulo:           input.Titulo,
		Genero:           input.Genero,
		FechaPublicacion: input.FechaPublicacion,
		EditorialID:      input.EditorialID,
		AutoresId:        input.AutoresId,
	}

	err = app.models.Book.CreateBook(book)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusAccepted, envelope{"book": book}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) updateBookHandler(w http.ResponseWriter, r *http.Request) {

	bookID := r.PathValue("id")

	if bookID == "" {
		app.notProvidedID(w, r)
		return
	}

	var input data.Book

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	book := &data.Book{
		IdLibro:          input.IdLibro,
		Titulo:           input.Titulo,
		Genero:           input.Genero,
		FechaPublicacion: input.FechaPublicacion,
		EditorialID:      input.EditorialID,
		AutoresId:        input.AutoresId,
	}

	err = app.models.Book.UpdateBook(book)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusAccepted, envelope{"book": book}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}

func (app *application) getFilteredBooksHandler(w http.ResponseWriter, r *http.Request) {

	estado := r.URL.Query().Get("estado")
	editorial := r.URL.Query().Get("editorial")

	books, err := app.models.Book.GetFilteredBooks(estado, editorial)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"books": books}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getBooksByGenreAndAuthorHandler(w http.ResponseWriter, r *http.Request) {

	genero := r.URL.Query().Get("genero")
	autor := r.URL.Query().Get("autor")

	books, err := app.models.Book.GetBooksByGenreAndAuthor(genero, autor)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	if len(books) == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"message": "No hay libros disponibles para el criterio especificado"})
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"books": books}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getBooksByPublicationDateHandler(w http.ResponseWriter, r *http.Request) {

	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")

	if startDate == "" || endDate == "" {
		http.Error(w, "Los parámetros 'start_date' y 'end_date' son obligatorios", http.StatusBadRequest)
		return
	}

	
	books , err := app.models.Book.GetBooksByPublicationDate(startDate , endDate)
	if err != nil{
		app.serverErrorResponse(w , r , err )
	}

	if len(books) == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"message": "No hay libros publicados en el rango de fechas especificado"})
		return
	}

	
	err = app.writeJSON(w, http.StatusOK, envelope{"books": books}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}

func (app *application) getBooksAvailableByGenreAndAuthorHandler(w http.ResponseWriter, r *http.Request) {


	genero := r.URL.Query().Get("genero")
	autor := r.URL.Query().Get("autor")
	titulo := r.URL.Query().Get("titulo")

	
	books , err := app.models.Book.GetBooksAvailableByGenreAndAuthor( genero , autor , titulo)
	if err != nil{
		app.serverErrorResponse(w , r , err )
	}

	if len(books) == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"message": "No hay libros disponibles para el criterio especificado"})
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"books": books}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}

func (app *application) getUnavailableBooksHandler(w http.ResponseWriter, r *http.Request) {

	books , err := app.models.Book.GetUnavailableBooks()
	
	w.Header().Set("Content-Type", "application/json")
	if len(books) == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"books": books}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
