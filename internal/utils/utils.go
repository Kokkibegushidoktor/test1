package utils

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/tech/closer"
)

const gracefulShutdownWaitTime = 2 * time.Second

func GracefulShutdown() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(ch)

	sig := <-ch
	log.Info().Msgf("%s %v - %s", "Received shutdown signal:", sig, "Graceful shutdown started")

	closer.CloseAll()

	time.Sleep(gracefulShutdownWaitTime)
	log.Info().Msg("Graceful shutdown is done")
}
