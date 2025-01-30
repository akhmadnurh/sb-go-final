package services

import (
	"OneTix/configs"
	"OneTix/constants"
	"OneTix/models"
	"OneTix/utils"
	"OneTix/validators"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(body validators.RegisterCustomerBody) (status int, response gin.H) {

	var existingUser models.MstUser
	if err := configs.DB.Where("email = ? and deleted_at is null", body.Email).First(&existingUser).Error; err == nil {
		if existingUser.ID != 0 {
			return http.StatusBadRequest, utils.GenerateResponse(false, "User with this email already exists")
		}
	}

	body.Password, _ = utils.HashPassword(body.Password)

	var newUser = models.MstUser{
		Name:     body.Name,
		Email:    body.Email,
		Password: body.Password,
		Role:     constants.CUSTOMER,
	}

	if err := configs.DB.Create(&newUser).Error; err != nil {
		return http.StatusInternalServerError, utils.GenerateResponse(false, "Failed to create user")
	}

	return http.StatusOK, utils.GenerateResponse(true, "Register success")
}

func RegisterOrganizer(body validators.RegisterOrganizerBody) (status int, response gin.H) {
	var existingUser models.MstUser
	if err := configs.DB.Where("email = ? and deleted_at is null", body.Email).First(&existingUser).Error; err == nil {
		if existingUser.ID != 0 {
			return http.StatusBadRequest, utils.GenerateResponse(false, "User with this email already exists")
		}
	}

	body.Password, _ = utils.HashPassword(body.Password)

	var newUser = models.MstUser{
		Name:     body.Name,
		Email:    body.Email,
		Password: body.Password,
		Role:     constants.ORGANIZER,
	}

	if err := configs.DB.Create(&newUser).Error; err != nil {
		return http.StatusInternalServerError, utils.GenerateResponse(false, "Failed to create user")
	}

	newOrganizer := models.MstOrganizer{
		UserID:      newUser.ID,
		CompanyName: body.CompanyName,
		ContactInfo: body.ContactInfo,
	}

	if err := configs.DB.Create(&newOrganizer).Error; err != nil {
		return http.StatusInternalServerError, utils.GenerateResponse(false, "Failed to create organizer")
	}

	return http.StatusOK, utils.GenerateResponse(true, "Register success")
}
