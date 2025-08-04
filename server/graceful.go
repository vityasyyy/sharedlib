package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/vityasyyy/sharedlib/logger"
)

func RunGracefully(port string, handler http.Handler, db *sqlx.DB) {
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}

	go func() {
		logger.Log.Info().Msg("Server starting on port " + port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Log.Fatal().Err(err).Msg("Server failed to start")
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	defer stop()

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		logger.Log.Error().Err(err).Msg("Server shutdown error")
	}

	if err := db.Close(); err != nil {
		logger.Log.Error().Err(err).Msg("DB close error")
	}

	logger.Log.Info().Msg("Shutdown complete")
}
