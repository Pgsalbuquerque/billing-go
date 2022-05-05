package routes

import (
	"strateegy/billing-service/controller"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		billing := main.Group("billing")
		{
			billing.POST("/", controller.ChangePlan)
		}
	}

	return router
}
