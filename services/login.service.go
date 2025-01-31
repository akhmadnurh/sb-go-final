package services

import (
	"OneTix/configs"
	"OneTix/models"
	"OneTix/utils"
	"OneTix/validators"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(body validators.LoginBody) (status int, response gin.H) {

	var user models.MstUser
	if err := configs.DB.Where("email = ? and deleted_at is null", body.Email).First(&user).Error; err != nil {
		return http.StatusUnauthorized, utils.GenerateResponse(false, "Invalid email or password")
	}

	verifyPassword := utils.VerifyPassword(user.Password, body.Password)
	if !verifyPassword {
		return http.StatusUnauthorized, utils.GenerateResponse(false, "Invalid email or password")
	}

	tokenString, err := utils.GenerateAccessToken(user)

	if err != nil {
		return http.StatusInternalServerError, utils.GenerateResponse(false, "Failed to generate token")
	}

	return http.StatusOK, utils.GenerateResponse(true, "Login success", gin.H{
		"access_token": tokenString,
	})
}
