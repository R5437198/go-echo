package sex_type

import (
	"github.com/labstack/echo/v4"
	"go-echo/internal/application/sex_type"
	"go-echo/internal/presentation/handler/master_data/sex_type/gen"
	"net/http"
)

type Handler struct {
	App sex_type.Application
}

func New(app sex_type.Application) *Handler {
	return &Handler{App: app}
}

func (h *Handler) GetSexTypes(ctx echo.Context) error {
	sts, err := h.App.FetchAll()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &gen.ErrorResponse{
			Code:    "000",
			Message: err.Error(),
		})
	}

	var res []gen.SexTypeResponse
	for i, v := range sts {
		res[i] = gen.SexTypeResponse{
			Id:        v.Id.String(),
			Value:     v.Value,
			CreatedAt: v.CreatedAt.String(),
			UpdatedAt: v.UpdatedAt.String(),
		}
	}

	return ctx.JSON(http.StatusOK, res)
}
