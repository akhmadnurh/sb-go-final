package services

import (
	"OneTix/configs"
	"OneTix/models"
	"OneTix/utils"
	"OneTix/validators"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(body validators.RegisterCustomerBody) (status int, response gin.H) {

	var existingUser models.User
	if err := configs.DB.Where("email = ? and deleted_at is null", body.Email).First(&existingUser).Error; err == nil {
		if existingUser.ID != 0 {
			return http.StatusBadRequest, utils.GenerateResponse(false, "User with this email already exists")
		}
	}

	body.Password, _ = utils.HashPassword(body.Password)

	var newUser = models.User{
		Name:     body.Name,
		Email:    body.Email,
		Password: body.Password,
		Role:     "customer",
	}

	if err := configs.DB.Create(&newUser).Error; err != nil {
		return http.StatusInternalServerError, utils.GenerateResponse(false, "Failed to create user")
	}

	return http.StatusOK, utils.GenerateResponse(true, "Register success")
}
