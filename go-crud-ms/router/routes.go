package router

import (
	"github.com/felipematheus1337/InsightFlow-AI/go-crud-ms/handler"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine, e *handler.EmployeeHandler) {

	basePath := "/api/v1/employees"

	v1 := router.Group(basePath)

	RegisterEmployeeRoutes(v1, e)
}

func RegisterEmployeeRoutes(v1 *gin.RouterGroup, e *handler.EmployeeHandler) {

	{
		v1.POST("/", e.CreateEmployee)
		v1.GET("/", e.GetEmployees)
	}
}
