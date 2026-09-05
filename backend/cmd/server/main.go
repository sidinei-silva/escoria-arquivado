package main

import (
	"context"
	"escoria/internal/account"
	"escoria/internal/gamedata"
	"escoria/internal/network/http"
	"escoria/internal/persistence/postgres"
	"log"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {
	slog.Info("Starting server...")

	if err := godotenv.Load(); err != nil {
		log.Println("warning: .env not found")
	}

	// Variáveis de ambiente
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL não definido")
	}

	// Gamedata
	var dataPathDir = "../data/"

	absolutePath, err := filepath.Abs(dataPathDir)
	if err != nil {
		slog.Error(
			"Falha ao obter caminho absoluto do arquivo",
			"dataPathDir", dataPathDir,
			"err", err,
		)
		log.Fatal(err)
	}

	slog.Info("Loading gamedata...")

	gd, err := gamedata.Load(absolutePath)
	if err != nil {
		slog.Error(
			"Falha ao carregar gamedata",
			"error", err,
		)
		log.Fatal(err)
	}

	slog.Info("Conteúdo carregado", "zones", len(gd.Zones))

	// Database
	db, err := pgx.Connect(
		context.Background(),
		databaseURL,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close(context.Background())

	// Dependencies
	accountRepository := postgres.NewAccountRepository(db)
	accountService := account.NewService(accountRepository)

	// HTTP
	accountHandler := http.NewAccountHandler(accountService)
	httpServer := http.NewServer(accountHandler)

	// Server
	slog.Info("HTTP listening on :8080")

	if err := httpServer.Start(); err != nil {
		log.Fatal(err)
	}
}
