package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/lmittmann/tint"
)

type application struct {
	logger *slog.Logger
}

func main() {
	port := flag.Int("port", 8080, "Port number to serve the server")
	flag.Parse()

	logger := slog.New(tint.NewHandler(os.Stdout, &tint.Options{
		Level:      slog.LevelInfo,
		TimeFormat: time.Kitchen,
	}))

	app := application{
		logger: logger,
	}

	router := http.NewServeMux()

	router.HandleFunc("GET /static/", func(w http.ResponseWriter, r *http.Request) {
		filePath := r.URL.Path[len("/static/"):]
		fullPath := filepath.Join(".", "static", filePath)
		http.ServeFile(w, r, fullPath)
	})

	router.HandleFunc("GET /healthcheck", app.healthcheckHandler)
	router.HandleFunc("GET /", app.homeHandler)
	router.HandleFunc("GET /mainInfo", app.mainInfoHandler)
	router.HandleFunc("GET /explanation", app.explanationHandler)

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", *port),
		Handler: router,
	}

	app.logger.Info(fmt.Sprintf("Starting server on port %s", server.Addr))
	server.ListenAndServe()
}
