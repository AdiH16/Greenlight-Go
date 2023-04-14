package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// App version
const version = "1.0.0"

// Configuration settings
type config struct {
	port int
	env  string
}

// Struct dependencies for HTTP handlers, helpers and middleware
type application struct {
	config config
	logger *log.Logger
}

func main() {
	//Struct instance
	var cfg config

	//Read the value of port and env command-line flags into the config struct
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	//Logger initialization
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	//Application struct
	app := &application{
		config: cfg,
		logger: logger,
	}

	//HTTP server with timeout settings
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	//Starting the HTTP server
	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)

}
