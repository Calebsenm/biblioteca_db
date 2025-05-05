package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) getPendingFinesHandler(w http.ResponseWriter, r *http.Request) {

	fine , err := app.models.Fine.GetPendingFines() 
	if err != nil {
		app.badRequestResponse(w , r , err)
		return 
	}

	err = app.writeJSON(w ,http.StatusOK , envelope{"fine": fine} , nil)
	if err != nil {
		app.serverErrorResponse( w , r , err)
	}
}


func (app *application) getUserFinesHandler(w http.ResponseWriter, r *http.Request) {

	idsocio := r.URL.Query().Get("idsocio")
	if idsocio == "" {
		http.Error(w, "El parámetro 'idsocio' es requerido", http.StatusBadRequest)
		return
	}

	userfines , err := app.models.Fine.GetUserFines(idsocio)
	if err != nil {
		app.badRequestResponse( w , r , err)
	}

	err = app.writeJSON(w , http.StatusOK, envelope{"userfines": userfines}, nil )
	if err != nil{
		app.serverErrorResponse(w , r , err )
	}
}


func (app *application) getUserPendingFinesHandler(w http.ResponseWriter, r *http.Request) {

	usuarioID := r.URL.Query().Get("usuario_id")
	if usuarioID == "" {
		http.Error(w, "El parámetro 'usuario_id' es obligatorio", http.StatusBadRequest)
		return
	}

	fines  , err := app.models.Fine.GetUserPendingFines(usuarioID)

	if len(fines) == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"message": "No hay multas pendientes para este usuario"})
		return
	}

	err = app.writeJSON(w , http.StatusOK, envelope{"userfines": fines}, nil )
	if err != nil{
		app.serverErrorResponse(w , r , err )
	}
}
