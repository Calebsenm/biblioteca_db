package main

import (
	"github.com/justinas/alice"
	"github.com/rs/cors"
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("GET /api/admin/books", app.getFilteredBooksHandler)
	mux.HandleFunc("GET /api/admin/books/unavailable", app.getUnavailableBooksHandler)
	mux.HandleFunc("GET /api/admin/users", app.getUsersByTypeHandler)
	mux.HandleFunc("GET /api/admin/loans", app.getActiveLoansHandler)
	mux.HandleFunc("GET /api/admin/fines/to", app.getPendingFinesHandler)
	mux.HandleFunc("GET /api/admin/fines", app.getUserFinesHandler)
	mux.HandleFunc("GET /api/admin/reservations", app.getActiveReservationsHandler)
	mux.HandleFunc("GET /api/admin/loans/history", app.getUserLoanHistoryHandler)
	mux.HandleFunc("GET /api/admin/books/available", app.getBooksByGenreAndAuthorHandler)
	mux.HandleFunc("GET /api/admin/books/published", app.getBooksByPublicationDateHandler)

	// Rutas para Usuario
	mux.HandleFunc("GET /api/books", app.getBooksAvailableByGenreAndAuthorHandler)
	mux.HandleFunc("POST /api/loans", app.createLoanHandler)
	mux.HandleFunc("GET /api/loans", app.getUserActiveLoanStatusHandler)
	mux.HandleFunc("GET /api/loans/completed", app.getUserCompletedLoanHistoryHandler)
	mux.HandleFunc("GET /api/fines", app.getUserPendingFinesHandler)
	mux.HandleFunc("GET /api/reservations", app.getUserActiveReservationsHandler) // GET

	mux.HandleFunc("POST /api/login", app.loginHandler)
	mux.HandleFunc("POST /api/register", app.registerHandler)

	// Rutas de Gestión de Libros
	mux.HandleFunc("POST /api/admin/books", app.createBookHandler)
	mux.HandleFunc("PUT /api/admin/books/{id}", app.updateBookHandler)
	mux.HandleFunc("PUT /api/editoriales", app.createEditorialHandler)
	mux.HandleFunc("GET /api/editoriales", app.getEditorialsHandler)
	mux.HandleFunc("GET /api/autores", app.getAutoresHandler)

	// Rutas de Gestión de Reservas
	mux.HandleFunc("POST /api/reservation", app.createReservation)
	mux.HandleFunc("DELETE /api/reservations/{id}", app.cancelReservationHandler)

	// Rutas de Gestión de Préstamos
	mux.HandleFunc("POST /api/loans/extend/{id}", app.extendLoanHandler) // POST - Extender préstamo

	c := cors.New(cors.Options{
		//AllowedOrigins:   []string{"http://127.0.0.1:3000"}, // Dominios permitidos
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders , c.Handler)

	return standard.Then(mux)
}
