package routes

import (
	"github.com/falasefemi2/hr-backend/controllers"
	"github.com/gin-gonic/gin"
)

func OnboardingRoutes(router *gin.Engine) {
	onboardingGroup := router.Group("/onboarding-workflow")
	{
		onboardingGroup.POST("/", controllers.CreateOnboardingWorkflow)
		onboardingGroup.GET("/all", controllers.GetAllOnboardingWorkflows)
		onboardingGroup.GET("/:id", controllers.GetSingleOnboardingWorkflow)
		onboardingGroup.POST("/assign", controllers.AssignWorkflowToEmployee)
	}
}
