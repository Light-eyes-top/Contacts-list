package hendler_REST

import (
	"SomeRestApi/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	contact := router.Group("/contact")
	{
		contact.POST("/post", h.CreateContact)
		contact.GET("/get/:id", h.GetContact)
		contact.PUT("/put/:id", h.UpdateContact)
		contact.DELETE("/delete/:id", h.DeleteContact)
	}
	return router
}
