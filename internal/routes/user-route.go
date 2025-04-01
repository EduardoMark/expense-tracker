package routes

import (
	"github.com/EduardoMark/expense-tracker/internal/handlers"
	"github.com/gin-gonic/gin"
)

func userRoutes(r *gin.Engine) {
	routes := r.Group("/users")
	routes.POST("", handlers.Create)
	routes.GET("")
	routes.GET("/:id")
	routes.PUT("/:id")
	routes.DELETE("/:id")
}
