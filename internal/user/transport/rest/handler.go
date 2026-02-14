package grpc

import (
	"net/http"

	"github.com/abdillahfazri/grpc-simple/internal/user/usecase"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	uc usecase.UserUsecase
}

func NewHandler(uc usecase.UserUsecase) *Handler {
	return &Handler{uc: uc}
}

func (h *Handler) GetUser(c *gin.Context) {
	id := c.Param("id")

	user, err := h.uc.GetUser(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
	})
}
