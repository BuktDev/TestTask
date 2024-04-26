package app

import (
	"github.com/labstack/echo/v4"
	"log"
	"middleware/internal/app/service"

	"middleware/internal/app/endpoint"
	"middleware/internal/app/mw"
)

type App struct {
	e    *endpoint.Endpoint
	s    *service.Service
	echo *echo.Echo
}

func New() (*App, error) {
	a := &App{}

	a.s = service.New()

	a.e = endpoint.New(a.s)

	a.echo = echo.New()

	a.echo.Use(mw.RoleCheck)

	a.echo.GET("/status", a.e.Status)

	return a, nil
}

func (a *App) Start() error {

	err := a.echo.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
