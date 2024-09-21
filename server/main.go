package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/lib/pq" // Import driver PostgreSQL
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"server/gomovie/app"
	"server/gomovie/controller"
	"server/gomovie/repository"
	"server/gomovie/service"
)

type config struct {
	port int
	db   struct {
		dsn string
	}
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4001, "Server port to listen on")
	flag.StringVar(&cfg.db.dsn, "dsn", "postgres://postgres:root123@localhost/id_gomoviereact?sslmode=disable", "Postgres connection config")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := openDB(cfg)
	if err != nil {
		logger.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(5) // Kurangi jumlah koneksi maksimal
	db.SetMaxIdleConns(5)

	validate := validator.New()
	logger.Println("Initializing movie repository...")
	movieRepo := repository.NewMovieRepository(db)

	logger.Println("Initializing movie service...")
	movieService := service.NewMovieService(movieRepo, db, validate)

	logger.Println("Initializing movie controller...")
	movieController := controller.NewMovieController(movieService)

	router := app.NewRouter(movieController)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("Starting server on port %d", cfg.port)
	err = server.ListenAndServe()
	if err != nil {
		logger.Fatalf("Server error: %v", err)
	}
}

// openDB opens a connection to the database and verifies it is reachable.
func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, fmt.Errorf("unable to open database connection: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to ping database: %w", err)
	}

	log.Println("Database connection established successfully")
	return db, nil
}
