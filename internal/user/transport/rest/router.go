package grpc

import (
	"github.com/gin-gonic/gin"
)

	func RegisterRoutes(
	r *gin.Engine,
	h *Handler,
) {
	r.GET("/users/:id", h.GetUser)
}
