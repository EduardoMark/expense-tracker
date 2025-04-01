package routes

import "github.com/gin-gonic/gin"

func SetupServer() {
	router := gin.Default()

	// user routes
	userRoutes(router)

	router.Run(":3000")
}
