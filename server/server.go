package server

import (
	"dungeons-and-dragons/db"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type Server struct {
	Echo *echo.Echo
	Db   *gorm.DB
}

func NewServer() *Server {
	s := &Server{
		Echo: echo.New(),
		Db:   db.Init(),
	}

	return s
}

func (s *Server) Start(addr string) error {
	return s.Echo.Start(":" + addr)
}
