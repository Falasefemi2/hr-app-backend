package routes

import (
	"github.com/falasefemi2/hr-backend/controllers"
	"github.com/gin-gonic/gin"
)

func CompanyRoutes(router *gin.Engine) {
	companyGroup := router.Group("/companies")
	{
		companyGroup.POST("/register", controllers.RegisterCompany)
		companyGroup.POST("/login", controllers.LoginCompany)
		companyGroup.GET("/all", controllers.GetAllCompanies)
		companyGroup.GET("/:id", controllers.GetSingleCompany)
	}
}
