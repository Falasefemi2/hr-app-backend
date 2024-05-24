package controllers

import "github.com/gin-gonic/gin"

func CreateOnboardingWorkflow(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Create Onboarding Workflow",
	})
}

func GetAllOnboardingWorkflows(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get All Onboarding Workflows",
	})
}

func GetSingleOnboardingWorkflow(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get Single Onboarding Workflow",
	})
}

func AssignWorkflowToEmployee(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Assign Workflow To Employee",
	})
}
