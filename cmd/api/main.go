package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/gaba-bouliva/movent/internal/data"
	"github.com/joho/godotenv"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *slog.Logger
	db     *data.Queries
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment [development|staging|production]")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	err := godotenv.Load()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	dsn := os.Getenv("DB_URL")
	if dsn == "" {
		logger.Error("invalid data source name")
		os.Exit(1)
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	dbQueries := data.New(db)

	app := &application{
		config: cfg,
		logger: logger,
		db:     dbQueries,
	}

	err = app.run()
	logger.Error(err.Error())
	os.Exit(1)

}

func (app *application) run() error {
	server := http.Server{
		Addr:        fmt.Sprintf(":%d", app.config.port),
		Handler:     app.router(),
		IdleTimeout: 10 * time.Second,
		ReadTimeout: 10 * time.Second,
		ErrorLog:    slog.NewLogLogger(app.logger.Handler(), slog.LevelError),
	}
	app.logger.Info(fmt.Sprintf("starting %s server on port %d", app.config.env, app.config.port))
	return server.ListenAndServe()
}
