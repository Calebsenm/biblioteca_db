package data

import (
	"context"
	"database/sql"
	"time"
)

type Book struct {
	IdLibro          int    `json:"idlibro"`
	Titulo           string `json:"titulo"`
	Genero           string `json:"genero"`
	FechaPublicacion string `json:"fechapublicacion"`
	EditorialID      int    `json:"ideditorial"`
	AutoresId        []int  `json:"idautores"`
	Status           string
}

type BookResponseGenAut struct {
	IDLibro int    `json:"id_libro"`
	Titulo  string `json:"titulo"`
	Genero  string `json:"genero"`
	Estado  string `json:"estado"`
	Autor   string `json:"autor"`
}

type BookbyDate struct {
	IDLibro          int    `json:"id_libro"`
	Titulo           string `json:"titulo"`
	Genero           string `json:"genero"`
	FechaPublicacion string `json:"fecha_publicacion"`
	Estado           string `json:"estado"`
	Editorial        string `json:"editorial"`
}

type BookAbailable struct {
	IDLibro int    `json:"id_libro"`
	Titulo  string `json:"titulo"`
	Genero  string `json:"genero"`
	Estado  string `json:"estado"`
	Autor   string `json:"autor"`
}

type BookUnavailable struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Genre  string `json:"genre"`
	Status string `json:"status"`
}

type BookModel struct {
	DB *sql.DB
}

func (m BookModel) CreateBook(book *Book) error {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	sqlQuery := `INSERT INTO libro (ideditorial, fechapublicacion,titulo, genero, estado)
        VALUES (?, ?, ?, ?,'disponible')`

	args := []any{book.EditorialID, book.FechaPublicacion, book.Titulo, book.Genero}

	result, err := m.DB.ExecContext(ctx, sqlQuery, args...)

	if err != nil {
		return err
	}

	bookID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	for _, autorID := range book.AutoresId {
		insertAuthorQuery := `
            INSERT INTO libro_autor (idlibro, idautor)
            VALUES (?, ?)`

		_, err = m.DB.Exec(insertAuthorQuery, bookID, autorID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m BookModel) UpdateBook(updatedBook *Book) error {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var existingBookCount int
	checkBookQuery := `SELECT COUNT(*) FROM libro WHERE idlibro = ?`

	bookId := updatedBook.IdLibro

	if err := m.DB.QueryRow(checkBookQuery, bookId).Scan(&existingBookCount); err != nil || existingBookCount == 0 {
		return err
	}

	updateBookQuery := `
			UPDATE libro
			SET titulo = ?, genero = ?, fechapublicacion = ?, ideditorial = ?
			WHERE idlibro = ?`

	args := []any{updatedBook.Titulo, updatedBook.Genero, updatedBook.FechaPublicacion, updatedBook.EditorialID, updatedBook.IdLibro}
	_, err := m.DB.ExecContext(ctx, updateBookQuery, args...)

	if err != nil {
		return err
	}

	deleteAuthorsQuery := `DELETE FROM libro_autor WHERE idlibro = ?`

	_, err = m.DB.Exec(deleteAuthorsQuery, bookId)

	if err != nil {
		return err
	}

	for _, autorID := range updatedBook.AutoresId {
		insertAuthorQuery := `INSERT INTO libro_autor (idlibro, idautor) VALUES (?, ?)`
		_, err := m.DB.Exec(insertAuthorQuery, bookId, autorID)
		if err != nil {

			return err
		}
	}

	return nil
}

func (m BookModel) GetFilteredBooks(estado string, editorial string) ([]*Book, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	querySQL := `
		SELECT 
			libro.idlibro, libro.titulo, libro.genero, libro.estado, editorial.nombre AS editorial
		FROM 
			libro
		INNER JOIN 
			editorial ON libro.ideditorial = editorial.ideditorial
		WHERE 
			(estado = ? OR ? = '') AND 
			(editorial.nombre = ? OR ? = '')`

	rows, err := m.DB.QueryContext(ctx, querySQL, estado, estado, editorial, editorial)
	if err != nil {
		return nil, err
	}

	var books []*Book

	for rows.Next() {
		var book Book

		err := rows.Scan(&book.IdLibro, &book.Titulo, &book.Genero, &book.Status, &book.EditorialID)
		if err == nil {
			books = append(books, &book)
		}
	}
	return books, nil
}

func (m BookModel) GetBooksByGenreAndAuthor(genero string, autor string) ([]*BookResponseGenAut, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
	SELECT 
		libro.idlibro,
		libro.titulo,
		libro.genero,
		libro.estado,
		autor.nombre AS autor
	FROM 
		libro
	JOIN 
		libro_autor ON libro.idlibro = libro_autor.idlibro
	JOIN 
		autor ON libro_autor.idautor = autor.idautor
	WHERE 
		libro.estado = 'disponible'
		AND libro.genero = ?
		AND autor.nombre LIKE ?
	`

	rows, err := m.DB.QueryContext(ctx, query, genero, "%"+autor+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []*BookResponseGenAut

	for rows.Next() {
		var book BookResponseGenAut
		err := rows.Scan(&book.IDLibro, &book.Titulo, &book.Genero, &book.Estado, &book.Autor)
		if err != nil {

			return nil, err
		}
		books = append(books, &book)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func (m BookModel) GetBooksByPublicationDate(startDate, endDate string) ([]*BookbyDate, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
			SELECT 
				libro.idlibro,
				libro.titulo,
				libro.genero,
				libro.fechapublicacion,
				libro.estado,
				editorial.nombre AS editorial
			FROM 
				libro
			JOIN 
				editorial ON libro.ideditorial = editorial.ideditorial
			WHERE 
				libro.fechapublicacion BETWEEN ? AND ?
			`
	rows, err := m.DB.QueryContext(ctx, query, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []*BookbyDate

	for rows.Next() {
		var book BookbyDate
		err := rows.Scan(&book.IDLibro, &book.Titulo, &book.Genero, &book.FechaPublicacion, &book.Estado, &book.Editorial)
		if err != nil {
			return nil, err
		}
		books = append(books, &book)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return books, nil
}

func (m BookModel) GetBooksAvailableByGenreAndAuthor(genero , autor , titulo string) ([]*BookAbailable , error) {
	
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	
	query := `
        SELECT 
            libro.idlibro,
            libro.titulo,
            libro.genero,
            libro.estado,
            autor.nombre AS autor
        FROM 
            libro
        JOIN 
            libro_autor ON libro.idlibro = libro_autor.idlibro
        JOIN 
            autor ON libro_autor.idautor = autor.idautor
        WHERE 
            libro.estado = 'disponible' 
    	`

	var params []interface{}

	if genero != "" {
		query += " AND libro.genero LIKE ?"
		params = append(params, "%"+genero+"%")
	}
	if autor != "" {
		query += " AND autor.nombre LIKE ?"
		params = append(params, "%"+autor+"%")
	}
	if titulo != "" {
		query += " AND libro.titulo LIKE ?"
		params = append(params, "%"+titulo+"%")
	}


	rows, err := m.DB.QueryContext(ctx , query, params...)
	if err != nil {
		return nil , err 
	}
	defer rows.Close()

	

	var books []*BookAbailable

	for rows.Next() {
		var book BookAbailable
		err := rows.Scan(&book.IDLibro, &book.Titulo, &book.Genero, &book.Estado, &book.Autor)
		if err != nil {
			return nil , err 
		}
		books = append(books, &book)
	}

	if err := rows.Err(); err != nil {
		return nil , err
	}

	return books , nil
}

func (m BookModel) GetUnavailableBooks() ([]*BookUnavailable , error) {
	
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
    SELECT 
        libro.idlibro, libro.titulo, libro.genero, libro.estado
    FROM 
        libro
    WHERE 
        libro.estado IN ('prestado', 'reservado')`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil , err
	}

	var books []*BookUnavailable

	for rows.Next() {
		var book BookUnavailable
		if err := rows.Scan(&book.ID, &book.Title, &book.Genre, &book.Status); err != nil {
			return nil , err 
		}
		books = append(books, &book)
	}

	if err := rows.Err(); err != nil {
		return nil , err 
	}

	return books , nil
}
