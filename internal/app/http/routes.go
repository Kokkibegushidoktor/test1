package http

func (s *Server) setupRoutes() {
	s.server.Post("/add_ticker", s.handler.AddTicker)
	s.server.Get("/fetch", s.handler.Fetch)
}
