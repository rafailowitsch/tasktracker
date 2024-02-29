package app

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"tasktracker/internal/config"
	delivery "tasktracker/internal/delivery/http"
	repository "tasktracker/internal/repository/postgres"
	"tasktracker/internal/service"
	"tasktracker/pkg/log/sl"
	"tasktracker/pkg/log/slogpretty"
	"time"

	"github.com/jackc/pgx/v5"
)

func Run(configsDir, env string) {
	cfg := config.MustLoad(configsDir, env)

	log := setupLogger(cfg.Env)

	log.Info("start app")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	connString := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s",
		cfg.Postgres.Username,
		cfg.Postgres.Password,
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.DBName,
	)

	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		log.Error("Unable to establish connection", sl.Err(err))
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	log.Info("database connect")

	repos := repository.NewRepositories(conn)
	service := service.NewServices(service.Deps{
		Repos: repos,
		Log:   *log,
	})
	handler := delivery.NewHandler(service, *log)
	var srv http.Server
	srv.Handler = handler.InitRoutes()
	srv.Addr = "localhost:8000"

	srv.ListenAndServe()
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case "local":
		log = setupPrettySlog()
	case "dev":
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case "prod":
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	default: // If env config is invalid, set prod settings by default due to security
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
