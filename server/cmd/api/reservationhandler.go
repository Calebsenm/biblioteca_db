package main

import (
	"biblioteca/internal/data"
	"encoding/json"
	"net/http"
)

func (app *application) createReservation(w http.ResponseWriter, r *http.Request) {

	var input data.Reservation

	err := app.readJSON(w, r, &input)
	if err != nil {
		http.Error(w, "Error al procesar la solicitud", http.StatusBadRequest)
		return
	}

	reservation := &data.Reservation{
		SocioID:      input.SocioID,
		LibroID:      input.LibroID,
		FechaReserva: input.FechaReserva,
	}

	err = app.models.Reservation.CreateReservation(reservation)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{"reservation": reservation}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) cancelReservationHandler(w http.ResponseWriter, r *http.Request) {

	reservationID := r.PathValue("id")

	if reservationID == "" {
		http.Error(w, "ID del reserva no proporcionado", http.StatusBadRequest)
		return
	}

	err := app.models.Reservation.CancelReservation(reservationID)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Reserva cancelada exitosamente"))
}

func (app *application) getActiveReservationsHandler(w http.ResponseWriter, r *http.Request) {
	
	usuario := r.URL.Query().Get("usuarioid")
	libro := r.URL.Query().Get("libro")
	fecha := r.URL.Query().Get("fecha")
	nombreSocio := r.URL.Query().Get("nombre")

	
	reservations , err := app.models.Reservation.GetActiveReservations(usuario , libro , fecha , nombreSocio)
	if err != nil {
		app.badRequestResponse(w , r , err )
		return 
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{"reservations": reservations}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
	
}


func (app *application) getUserActiveReservationsHandler(w http.ResponseWriter, r *http.Request) {


	usuarioID := r.URL.Query().Get("usuario_id")
	if usuarioID == "" {
		http.Error(w, "El parámetro 'usuario_id' es obligatorio", http.StatusBadRequest)
		return
	}

	reservations , err := app.models.Reservation.GetUserActiveReservations(usuarioID)
	if err != nil{
		app.badRequestResponse( w , r , err )
		return
	}

	if len(reservations) == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"message": "No hay reservas activas para este usuario"})
		return
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{"reservations": reservations}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
