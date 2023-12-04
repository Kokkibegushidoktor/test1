package http

func (s *Server) setupRoutes() {
	s.server.Get("/add_ticker", s.handler.AddTicker)
}
