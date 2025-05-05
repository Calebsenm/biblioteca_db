package main

import (
	"biblioteca/internal/data"
	"encoding/json"
	"net/http"
	"time"
)

func (app *application) createLoanHandler(w http.ResponseWriter, r *http.Request) {

	var input data.Loan

	err := app.readJSON(w, r, &input)
	if err != nil {
		http.Error(w, "Error al decodificar el cuerpo de la solicitud", http.StatusBadRequest)
		return
	}

	if input.UsuarioID == 0 || input.LibroID == 0 || input.FechaPrestamo == "" || input.FechaDevolucion == "" {
		http.Error(w, "Todos los parámetros son obligatorios", http.StatusBadRequest)
		return
	}

	loan := &data.Loan{
		LibroID:         input.LibroID,
		FechaPrestamo:   input.FechaPrestamo,
		FechaDevolucion: input.FechaDevolucion,
	}

	err = app.models.Loan.CreateLoan(loan)
	
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{"loan": loan}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}

func (app *application) extendLoanHandler(w http.ResponseWriter, r *http.Request) {

	reservationID := r.PathValue("id")
	
	if reservationID == "" {
		http.Error(w, "ID del reserva no proporcionado", http.StatusBadRequest)
		return
	}
	
	var request struct {
		NuevaFechaDevolucion string `json:"nuevafechadevolucion"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Error al procesar la solicitud", http.StatusBadRequest)
		return
	}

	nuevaFecha, err := time.Parse("2006-01-02", request.NuevaFechaDevolucion)
	if err != nil {
		http.Error(w, "Fecha inválida, el formato debe ser AAAA-MM-DD", http.StatusBadRequest)
		return
	}

	err = app.models.Loan.ExtendLoand(reservationID , nuevaFecha )
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Préstamo extendido exitosamente"))
}

func (app *application) getActiveLoansHandler(w http.ResponseWriter, r *http.Request) {

	userID := r.URL.Query().Get("idsocio")
	startDate := r.URL.Query().Get("startdate")
	endDate := r.URL.Query().Get("enddate")

	if userID == "" || startDate == "" || endDate == "" {
		http.Error(w, "idsocio, startdate, y enddate son requeridos", http.StatusBadRequest)
		return
	}

	loans , err := app.models.Loan.GetActiveLoans(userID , startDate , endDate )
	if err != nil {
		app.badRequestResponse(w, r, err)
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{"loansActive": loans}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getUserLoanHistoryHandler(w http.ResponseWriter, r *http.Request) {
	idUsuario := r.URL.Query().Get("idsocio")

	loans , err := app.models.Loan.GetUserLoanHistory(idUsuario)
	if err != nil {
		app.badRequestResponse(w, r, err)
	}

	if len(loans) == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"message": "No hay préstamos registrados para este usuario"})
		return
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{"loanshystory": loans}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getUserActiveLoanStatusHandler(w http.ResponseWriter, r *http.Request) {

	usuarioID := r.URL.Query().Get("usuario_id")
	if usuarioID == "" {
		http.Error(w, "El parámetro 'usuario_id' es obligatorio", http.StatusBadRequest)
		return
	}


	loans , err := app.models.Loan.GetUserActiveLoanStatus(usuarioID)
	if err != nil {
		app.badRequestResponse(w, r, err)
	}

	if len(loans) == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"message": "No hay préstamos activos para este usuario"})
		return
	}

	
	err = app.writeJSON(w, http.StatusCreated, envelope{"loansactive": loans}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}

func (app *application) getUserCompletedLoanHistoryHandler(w http.ResponseWriter, r *http.Request) {

	usuarioID := r.URL.Query().Get("usuario_id")
	if usuarioID == "" {
		http.Error(w, "El parámetro 'usuario_id' es obligatorio", http.StatusBadRequest)
		return
	}

	
	loans , err := app.models.Loan.GetUserCompletedLoanHistory(usuarioID)
	if err != nil {
		app.badRequestResponse(w, r, err)
	}

	if len(loans) == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"message": "No hay préstamos completados para este usuario"})
		return
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{"loanscomplete": loans}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
