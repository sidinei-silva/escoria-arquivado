package main

import (
	dataloader "escoria/internal/data_loader"
	"escoria/internal/network"
	"log"
	"log/slog"
	"net/http"
)

func main() {
	slog.Info("Starting server...")

	zones, err := dataloader.LoadZonesFromFile()
	if err != nil {
		log.Fatal(err)
	}

	slog.Info("Loaded zones", "count", len(zones))

	httpServer := network.NewHTTPServer()

	slog.Info("HTTP listening on :8080")

	if err := http.ListenAndServe(":8080", httpServer.Handler()); err != nil {
		log.Fatal(err)
	}
}
