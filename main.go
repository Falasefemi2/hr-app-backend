package main

import (
	"os"

	"github.com/falasefemi2/hr-backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}
	router := gin.New()
	router.Use(gin.Logger())

	routes.CompanyRoutes(router)
	routes.EmployeeRoutes(router)
	routes.OnboardingRoutes(router)

	// Start the server
	router.Run(":" + port)

}
