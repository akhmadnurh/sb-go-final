package controllers

import (
	"OneTix/services"
	"OneTix/structs"
	"OneTix/validators"

	"github.com/gin-gonic/gin"
)

func UpdateProfile(c *gin.Context) {
	requestBody, _ := c.Get("validated_body")
	body := requestBody.(validators.UpdateProfileBody)

	requestUser, _ := c.Get("user")
	user := requestUser.(structs.JWTClaims)

	status, response := services.UpdateProfile(user.ID, body)

	c.JSON(status, response)
}

func UpdateOrganizerProfile(c *gin.Context) {
	requestBody, _ := c.Get("validated_body")
	body := requestBody.(validators.UpdateOrganizerProfileBody)

	requestUser, _ := c.Get("user")
	user := requestUser.(structs.JWTClaims)

	status, response := services.UpdateOrganizerProfile(user.ID, body)

	c.JSON(status, response)
}

func GetProfile(c *gin.Context) {
	requestUser, _ := c.Get("user")
	user := requestUser.(structs.JWTClaims)

	status, response := services.GetProfile(user.ID)

	c.JSON(status, response)
}

func GetOrganizerProfile(c *gin.Context) {
	requestUser, _ := c.Get("user")
	user := requestUser.(structs.JWTClaims)

	status, response := services.GetOrganizerProfile(user.ID)

	c.JSON(status, response)
}

func UpdatePassword(c *gin.Context) {
	requestBody, _ := c.Get("validated_body")
	body := requestBody.(validators.UpdatePasswordBody)

	requestUser, _ := c.Get("user")
	user := requestUser.(structs.JWTClaims)

	status, response := services.UpdatePassword(user.ID, body)

	c.JSON(status, response)
}
