package main

import (
	"database/sql"
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
)

type config struct {
	port string
	env  string
	db   struct {
		dns string
	}
}

type application struct {
    errorLog *log.Logger
	infoLog  *log.Logger
	config   config
	db       *sql.DB
}

func main() {
	var con config

	flag.StringVar(&con.port, "addr", "4000", "HTTP network address port for API")
	flag.StringVar(&con.db.dns, "dns", "root:admin@(127.0.0.1:3306)/biblioteca?parseTime=true", "MySQL data source name")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime );
    errorLog := log.New(os.Stderr , "ERROR\t", log.Ldate | log.Ltime | log.Lshortfile );


	db, err := openDB(con)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	log.Println("Database connection pool established")

	app := &application{
		errorLog: errorLog,
		infoLog:infoLog,
		config: con,
		db:     db,
	}

	svr := &http.Server{
		Addr:    ":" + app.config.port,
		Handler: app.routes(),
	}

	err = svr.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

func openDB(con config) (*sql.DB, error) {
	db, err := sql.Open("mysql", con.db.dns)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
