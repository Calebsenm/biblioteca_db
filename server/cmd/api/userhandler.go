package main

import "net/http"

func (app *application) getUsersByTypeHandler(w http.ResponseWriter, r *http.Request) {
	userType := r.URL.Query().Get("tiposocio")

	if userType == "" {
		http.Error(w, "tipo de socio es requerido", http.StatusBadRequest)
		return
	}

	validUserTypes := map[string]bool{
		"normal":     true,
		"estudiante": true,
		"profesor":   true,
	}
	if !validUserTypes[userType] {
		http.Error(w, "tipo de socio no válido", http.StatusBadRequest)
		return
	}

	usertype, err := app.models.User_.GetUserByType(userType)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"usertype": usertype}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
