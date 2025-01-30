package controllers

import (
	"OneTix/services"
	"OneTix/structs"
	"OneTix/validators"

	"github.com/gin-gonic/gin"
)

func FindTickets(c *gin.Context) {
	requestQuery, _ := c.Get("validated_query")
	query := requestQuery.(validators.FindTicketsQuery)

	requestUser, _ := c.Get("user")
	user := requestUser.(structs.JWTClaims)

	status, response := services.FindTickets(user.ID, query)

	c.JSON(status, response)
}

func CreateTicket(c *gin.Context) {
	requestBody, _ := c.Get("validated_body")
	body := requestBody.(validators.CreateTicketBody)

	requestUser, _ := c.Get("user")
	user := requestUser.(structs.JWTClaims)

	status, response := services.CreateTicket(user.ID, body)

	c.JSON(status, response)
}

func CheckIn(c *gin.Context) {
	requestBody, _ := c.Get("validated_body")
	body := requestBody.(validators.CheckInBody)

	requestUser, _ := c.Get("user")
	user := requestUser.(structs.JWTClaims)

	status, response := services.CheckIn(user.ID, body)

	c.JSON(status, response)
}
