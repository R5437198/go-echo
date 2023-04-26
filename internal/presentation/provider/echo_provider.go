package provider

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-echo/internal/config"
	"go-echo/internal/presentation/router"
	"net/http"
)

func NewEcho() {
	c := config.New()
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	router.New(e)

	if err := e.Start(c.Port()); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			e.Logger.Fatal(err)
		}
	}
}
