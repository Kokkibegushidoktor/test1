package test1

import (
	"context"
	"github.com/Kokkibegushidoktor/test1/internal/app"
	"github.com/Kokkibegushidoktor/test1/internal/config"
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func main() {
	_ = godotenv.Load(".env")

	cfg := &config.Config{}

	if err := env.Parse(cfg); err != nil {
		log.Fatal().Msgf("Failed to load environment, err: %v", err)
	}

	if err := app.Run(context.Background(), cfg); err != nil {
		log.Fatal().Msgf("Error running service, err: %v", err)
	}
}
