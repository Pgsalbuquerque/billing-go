package controller

import (
	"strateegy/billing-service/controller/dto"
	"strateegy/billing-service/service"

	"github.com/gin-gonic/gin"
)

func ChangePlan(c *gin.Context) {
	var billing dto.Billing

	err := c.ShouldBindJSON(&billing)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "value is required",
		})
		return
	}

	token := c.GetHeader("Authorization")

	err = service.NewBillingService().ChangePlan(token, billing)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "value is required",
		})
		return
	}

	c.JSON(204, gin.H{})
}
