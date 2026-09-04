package main

import (
	"escoria/internal/gamedata"
	"escoria/internal/network"
	"log"
	"log/slog"
	"net/http"
	"path/filepath"
)

func main() {
	slog.Info("Starting server...")
	var dataPathDir = "../data/"
	absolutePath, err := filepath.Abs(dataPathDir)

	if err != nil {
		slog.Error("Falha ao obter caminho absoluto do arquivo", "zonePathFile", dataPathDir, "err", err)
		log.Fatal(err)
	}

	slog.Info("Loading gamedata...")
	gd, err := gamedata.Load(absolutePath)

	if err != nil {
		slog.Error("Falha ao carregar gamedata", "error", err)
		log.Fatal(err)
	}

	slog.Info("conteúdo carregado", "zones", len(gd.Zones))

	httpServer := network.NewHTTPServer()

	slog.Info("HTTP listening on :8080")

	if err := http.ListenAndServe(":8080", httpServer.Handler()); err != nil {
		log.Fatal(err)
	}
}
