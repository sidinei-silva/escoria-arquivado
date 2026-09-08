package bootstrap

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"escoria/internal/account"
	"escoria/internal/game"
	"escoria/internal/gamedata"
	httpnetwork "escoria/internal/network/http"
	"escoria/internal/persistence/postgres"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/lmittmann/tint"
)

type Application struct {
	httpServer *httpnetwork.Server
	db         *pgx.Conn
	world      *game.World
}

func New() (*Application, error) {
	initLogger()

	slog.Info("Starting server...")

	if err := godotenv.Load(); err != nil {
		log.Println("warning: .env not found")
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		return nil, fmt.Errorf("DATABASE_URL não definido")
	}

	dataPath, err := loadDataPath()
	if err != nil {
		return nil, err
	}

	data, err := gamedata.Load(dataPath)
	if err != nil {
		return nil, err
	}

	world, err := BuildWorld(data)
	if err != nil {
		return nil, err
	}

	slog.Info("Game world loaded", "zones", len(world.Zones))

	db, err := pgx.Connect(context.Background(), databaseURL)
	if err != nil {
		return nil, err
	}

	accountRepository := postgres.NewAccountRepository(db)
	accountService := account.NewService(accountRepository)

	accountHandler := httpnetwork.NewAccountHandler(accountService)
	httpServer := httpnetwork.NewServer(accountHandler)

	return &Application{
		httpServer: httpServer,
		db:         db,
		world:      world,
	}, nil
}

func (a *Application) Run() error {
	slog.Info("HTTP listening on :8080")

	return a.httpServer.Start()
}

func (a *Application) Shutdown(ctx context.Context) error {
	if err := a.httpServer.Shutdown(); err != nil {
		return err
	}

	return a.db.Close(ctx)
}

func loadDataPath() (string, error) {
	path, err := os.Getwd()

	if err != nil {
		return "", err
	}

	var dataPathDir = path + "/../data"

	absolutePath, err := filepath.Abs(dataPathDir)

	if err != nil {
		slog.Error(
			"Falha ao obter caminho absoluto do arquivo",
			"dataPathDir", dataPathDir,
			"err", err,
		)
		log.Fatal(err)
	}

	return absolutePath, nil
}

func initLogger() {
	w := os.Stderr
	logger := slog.New(tint.NewTextHandler(w, &tint.Options{
		Level:      slog.LevelDebug,
		TimeFormat: time.Kitchen,
		AddSource:  true,
	}))

	if os.Getenv("ENV") == "production" {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	}

	slog.SetDefault(logger)
}
