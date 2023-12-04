package http

import (
	"errors"
	"github.com/Kokkibegushidoktor/test1/internal/app/http/handlers"
	"github.com/Kokkibegushidoktor/test1/internal/config"
	"github.com/Kokkibegushidoktor/test1/internal/tech/closer"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"net/http"
)

type Server struct {
	server  *fiber.App
	handler *handlers.Handler
}

func New(cfg *config.Config, handler *handlers.Handler) *Server {
	server := fiber.New()

	return &Server{
		server:  server,
		handler: handler,
	}
}

func (s *Server) Start() {
	s.setupRoutes()

	go func() {
		err := s.server.Listen(":8080")
		if !errors.Is(err, http.ErrServerClosed) {
			log.Fatal().Err(err)
		}
	}()
	closer.Add(s.close)
}

func (s *Server) close() error {
	if err := s.server.Shutdown(); err != nil {
		log.Error().Msgf("Error stopping http server, err: %v", err)
		return err
	}

	log.Info().Msgf("Http server shutdown is done")

	return nil
}
