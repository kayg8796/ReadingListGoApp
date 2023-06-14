package main

import (
	"database/sql" //generic api for interacting with db in vendor neutral way
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"  //third party go package of the postgres db driver
	"readinglist.duffney.io/internal/data"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
	dsn string //db name service for interaction with db
}

type application struct { // consider that this type definition like a class and within it, its attributes. the methods are which ever function uses it as a receiver
	config config
	logger *log.Logger
	models data.Models
}

func main() { //entry point to the application
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "api server port")
	flag.StringVar(&cfg.env, "env", "dev", "Environment(dev|stage|prod)")
	flag.StringVar(&cfg.dsn, "db-dsn", os.Getenv("READINGLIST_DB_DSN"), "PostgreSQL DSN")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := sql.Open("postgres", cfg.dsn)   // open connection to the database
	if err != nil {
		logger.Fatal(err)
	}

	defer db.Close()

	err = db.Ping() //test connection to db
	if err != nil {
		logger.Fatal(err)
	}

	logger.Printf("database connection pool established")

	app := &application{
		config: cfg,
		logger: logger,
		models: data.NewModels(db),
	}

	addr := fmt.Sprintf(":%d", cfg.port)

	srv := &http.Server{
		Addr:         addr,
		Handler:      app.route(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting %s server on %s", cfg.env, addr)

	err = srv.ListenAndServe()
	logger.Fatal(err)
}
