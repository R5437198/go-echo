package router

import (
	"github.com/labstack/echo/v4"
	"go-echo/internal/presentation/controller/health_check"
	healthCheck "go-echo/internal/presentation/controller/health_check/gen"
)

func New(e *echo.Echo) {
	healthCheck.RegisterHandlers(e, health_check.NewServer())
}
