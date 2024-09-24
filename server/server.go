package server

import (
	"dnd-api/db"
	"dnd-api/db/stores"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type Server struct {
	Echo   *echo.Echo
	Db     *gorm.DB
	Stores *stores.Stores
}

func NewServer() *Server {
	s := &Server{
		Echo: echo.New(),
		Db:   db.Init(),
	}
	s.Stores = stores.NewStores(s.Db)

	return s
}

func (s *Server) Start(addr string) error {
	return s.Echo.Start(":" + addr)
}
