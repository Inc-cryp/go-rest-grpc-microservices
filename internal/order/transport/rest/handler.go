package rest

import (
	"errors"
	"net/http"

	"github.com/abdillahfazri/grpc-simple/internal/order/usecase"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	uc *usecase.OrderUsecase
}

func NewHandler(uc *usecase.OrderUsecase) *Handler {
	return &Handler{uc: uc}
}

func (h *Handler) CreateOrder(c *gin.Context) {
	userID := c.Param("user_id")

	result, err := h.uc.CreateOrder(c.Request.Context(), userID)
	if err != nil {
		c.JSON(mapHTTPStatus(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": result})
}

func mapHTTPStatus(err error) int {
	switch {
	case errors.Is(err, usecase.ErrUserNotFound):
		return http.StatusNotFound
	case errors.Is(err, usecase.ErrUserServiceTimeout):
		return http.StatusGatewayTimeout
	case errors.Is(err, usecase.ErrUserServiceUnavailable):
		return http.StatusServiceUnavailable
	default:
		return http.StatusBadGateway
	}
}
