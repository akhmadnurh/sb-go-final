package controllers

import (
	"OneTix/services"
	"OneTix/validators"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	requestBody, _ := c.Get("validated_body")
	body := requestBody.(validators.RegisterCustomerBody)

	status, response := services.Register(body)

	c.JSON(status, response)
}

func RegisterOrganizer(c *gin.Context) {
	requestBody, _ := c.Get("validated_body")
	body := requestBody.(validators.RegisterOrganizerBody)

	status, response := services.RegisterOrganizer(body)

	c.JSON(status, response)
}

func Login(c *gin.Context) {
	requestBody, _ := c.Get("validated_body")
	body := requestBody.(validators.LoginBody)

	status, response := services.Login(body)

	c.JSON(status, response)
}
