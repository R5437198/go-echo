package router

import (
	"github.com/labstack/echo/v4"
	"go-echo/internal/di"
	healthCheckGen "go-echo/internal/presentation/handler/health_check/gen"
	sexTypeGen "go-echo/internal/presentation/handler/master_data/sex_type/gen"
)

func New(e *echo.Echo) {
	healthCheckGen.RegisterHandlers(e, di.HealthCheckHandler())
	sexTypeGen.RegisterHandlers(e, di.SexTypeHandler())
}
