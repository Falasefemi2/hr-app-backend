package handlers

import "github.com/gin-gonic/gin"

func CreateEmployee(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Create employee",
	})
}

func LoginEmployee(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Login employee",
	})
}

func GetAllEmployees(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get all employees",
	})
}

func GetEmployeesByCompany(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get employees by company",
	})
}

func GetSingleEmployee(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get single employee",
	})
}

func EditEmployee(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Edit employee",
	})
}

func DeleteEmployee(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Delete employee",
	})
}
