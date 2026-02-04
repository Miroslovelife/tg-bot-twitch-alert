package server

import (
	"github.com/Miroslovelife/tg-bot-twitch-alert/internal/delivery/http/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type IServer interface {
}

type Server struct {
	Handlers []handler.Handler
}

func NewServer(handlers []handler.Handler) IServer {
	return &Server{
		Handlers: handlers,
	}
}

func (s *Server) MustInitServer() {
	e := echo.New()

	e.Use(middleware.RequestLogger())

	if err := e.Start(":1323"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
