package data

import (
	"context"
	"database/sql"
	"time"
)

type Fine struct {
	IDMulta    int     `json:"idmulta"`
	IDPrestamo int     `json:"idprestamo"`
	SaldoPagar float64 `json:"saldopagar"`
	FechaMulta string  `json:"fechamulta"`
	Estado     string  `json:"estado"`
}

type FineModel struct {
	DB *sql.DB
}

func (m FineModel) GetPendingFines() ([]*Fine, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		SELECT idmulta, idprestamo, saldopagar, fechamulta, estado
		FROM multa
		WHERE estado = 'pendiente'
	`
	rows, err := m.DB.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pendingFines []*Fine

	for rows.Next() {
		var fine Fine

		err := rows.Scan(&fine.IDMulta, &fine.IDPrestamo, &fine.SaldoPagar, &fine.FechaMulta, &fine.Estado)
		if err != nil {
			return nil, err
		}
		pendingFines = append(pendingFines, &fine)
	}
	return pendingFines, nil
}

func (m FineModel) GetUserFines(idsocio string) ([]*Fine, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		SELECT m.idmulta, m.idprestamo, m.saldopagar, m.fechamulta, m.estado
		FROM multa m
		INNER JOIN 
			prestamo p ON m.idprestamo = p.idprestamo
		WHERE p.idsocio = ?
		ORDER BY m.fechamulta DESC
	`
	rows, err := m.DB.QueryContext(ctx, query, idsocio)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var fines []*Fine

	for rows.Next() {
		var fine Fine
		err := rows.Scan(&fine.IDMulta, &fine.IDPrestamo, &fine.SaldoPagar, &fine.FechaMulta, &fine.Estado)
		if err != nil {
			return nil, err
		}
		fines = append(fines, &fine)
	}

	return fines, nil

}

func (m FineModel) GetUserPendingFines(usuario_id string) ([]*Fine, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		SELECT 
			multa.idmulta,
			multa.saldopagar,
			multa.fechamulta,
			multa.estado
		FROM 
			multa
		JOIN 
			prestamo ON multa.idprestamo = prestamo.idprestamo
		WHERE 
			prestamo.idsocio = ? AND multa.estado = 'pendiente'
	`

	rows, err := m.DB.QueryContext(ctx, query, usuario_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var fines []*Fine

	for rows.Next() {
		var fine Fine
		err := rows.Scan(&fine.IDMulta, &fine.SaldoPagar, &fine.FechaMulta, &fine.Estado)
		if err != nil {
			return nil, err
		}
		fines = append(fines, &fine)
	}

	if err := rows.Err(); err != nil {
		return nil , err
	}

	return fines , nil 
}
