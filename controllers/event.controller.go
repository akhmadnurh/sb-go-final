package controllers

import (
	"OneTix/services"
	"OneTix/structs"
	"OneTix/validators"

	"github.com/gin-gonic/gin"
)

func FindEvents(c *gin.Context) {
	requestQuery, _ := c.Get("validated_query")
	query := requestQuery.(validators.FindEventsQuery)

	status, response := services.FindEvents(query)

	c.JSON(status, response)
}

func CreateEvent(c *gin.Context) {
	requestUser, _ := c.Get("user")
	user := requestUser.(structs.JWTClaims)

	requestBody, _ := c.Get("validated_body")
	body := requestBody.(validators.CreateEventBody)

	status, response := services.CreateEvent(user.ID, body)

	c.JSON(status, response)
}

func DeleteEvent(c *gin.Context) {
	requestQuery, _ := c.Get("validated_query")
	query := requestQuery.(validators.DeleteEventQuery)

	status, response := services.DeleteEvent(query)

	c.JSON(status, response)
}

func UpdateEvent(c *gin.Context) {
	requestQuery, _ := c.Get("validated_query")
	query := requestQuery.(validators.UpdateEventQuery)

	requestBody, _ := c.Get("validated_body")
	body := requestBody.(validators.UpdateEventBody)

	status, response := services.UpdateEvent(query, body)

	c.JSON(status, response)
}
