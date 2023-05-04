package health_check

import (
	"github.com/labstack/echo/v4"
	"go-echo/internal/presentation/controller/health_check/gen"
	"net/http"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (Server) GetHealthCheck(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, &gen.DefaultResponse{Result: "ok"})
}
