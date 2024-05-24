package main

import (
	"github.com/falasefemi2/hr-backend/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// companies routes
	router.POST("/companies/register", handlers.RegisterCompany)
	router.POST("/companies/login", handlers.LoginCompany)
	router.GET("/companies/all", handlers.GetAllCompanies)
	router.GET("/companies/:id", handlers.GetSingleCompany)

	// Employee routes
	router.POST("/employees/create", handlers.CreateEmployee)
	router.POST("/employees/login", handlers.LoginEmployee)
	router.GET("/employees/all", handlers.GetAllEmployees)
	router.GET("/employees/all/:companyId", handlers.GetEmployeesByCompany)
	router.GET("/employees/:id", handlers.GetSingleEmployee)
	router.PUT("/employees/:id", handlers.EditEmployee)
	router.DELETE("/employees/:id", handlers.DeleteEmployee)

	// Onboarding Workflow routes
	router.POST("/onboarding-workflow", handlers.CreateOnboardingWorkflow)
	router.GET("/onboarding-workflow/all", handlers.GetAllOnboardingWorkflows)
	router.GET("/onboarding-workflow/:id", handlers.GetSingleOnboardingWorkflow)
	router.POST("/onboarding-workflow/assign", handlers.AssignWorkflowToEmployee)

	// Start the server
	router.Run(":8080")
}
