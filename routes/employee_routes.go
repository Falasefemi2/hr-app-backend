package routes

import (
	"github.com/falasefemi2/hr-backend/controllers"
	"github.com/gin-gonic/gin"
)

func EmployeeRoutes(router *gin.Engine) {
	employeeGroup := router.Group("/employees")
	{
		employeeGroup.POST("/create", controllers.CreateEmployee)
		employeeGroup.POST("/login", controllers.LoginEmployee)
		employeeGroup.GET("/all", controllers.GetAllEmployees)
		employeeGroup.GET("/all/:companyId", controllers.GetEmployeesByCompany)
		employeeGroup.GET("/:id", controllers.GetSingleEmployee)
		employeeGroup.PUT("/:id", controllers.EditEmployee)
		employeeGroup.DELETE("/:id", controllers.DeleteEmployee)
	}
}
