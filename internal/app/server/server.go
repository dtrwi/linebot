package server

import (
	"fmt"

	"github.com/dtrwi/linebot/internal/app/handler"
	"github.com/dtrwi/linebot/internal/app/service"
	"github.com/dtrwi/linebot/internal/pkg"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// IServer interface for server
type IServer interface {
	StartApp()
}

type server struct {
	pkg.Option
	*service.Service
}

// NewServer create object server
func NewServer(opt pkg.Option, svc *service.Service) IServer {
	return &server{
		Option:  opt,
		Service: svc,
	}
}

func (s *server) StartApp() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Pre(middleware.RemoveTrailingSlash())

	handler := handler.Option{
		Option:  s.Option,
		Service: s.Service,
	}
	Router(handler, e)

	address := fmt.Sprintf(":%d", s.Config.GetInt("app.port"))
	e.Logger.Fatal(e.Start(address))
}
