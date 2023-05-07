package health_check

import (
	"github.com/labstack/echo/v4"
	"go-echo/internal/presentation/handler/health_check/gen"
	"net/http"
)

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

func (Handler) GetHealthCheck(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, &gen.DefaultResponse{Result: "ok"})
}
