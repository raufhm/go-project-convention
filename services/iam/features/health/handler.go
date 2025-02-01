// services/iam/features/health/handler.go
package health

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	Service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) Register(e *echo.Echo) {
	if e == nil {
		panic("echo instance cannot be nil")
	}
	e.GET("/health", h.CheckHealth)
}

func (h *Handler) CheckHealth(c echo.Context) error {
	status := h.Service.Status()
	return c.JSON(http.StatusOK, status)
}
