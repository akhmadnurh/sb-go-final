package services

import (
	"OneTix/configs"
	"OneTix/models"
	"OneTix/utils"
	"OneTix/validators"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func UpdateProfile(userID uint, body validators.UpdateProfileBody) (status int, response gin.H) {
	var user models.MstUser

	if err := configs.DB.Where("id = ? and deleted_at is null", userID).First(&user).Error; err != nil {
		return http.StatusInternalServerError, utils.GenerateResponse(false, err.Error())
	}

	user.Name = body.Name
	user.UpdatedAt = time.Now()

	if err := configs.DB.Save(&user).Error; err != nil {
		return http.StatusInternalServerError, utils.GenerateResponse(false, err.Error())
	}

	return http.StatusOK, utils.GenerateResponse(true, "Success update profile")
}

func UpdateOrganizerProfile(userID uint, body validators.UpdateOrganizerProfileBody) (status int, response gin.H) {
	var user models.MstUser
	var organizer models.MstOrganizer

	if err := configs.DB.Where("id = ? and deleted_at is null", userID).First(&user).Error; err != nil {
		return http.StatusInternalServerError, utils.GenerateResponse(false, err.Error())
	}

	if err := configs.DB.Where("user_id = ? and deleted_at is null", userID).First(&organizer).Error; err != nil {
		return http.StatusInternalServerError, utils.GenerateResponse(false, err.Error())
	}

	user.Name = body.Name
	user.UpdatedAt = time.Now()
	organizer.CompanyName = body.CompanyName
	organizer.ContactInfo = body.ContactInfo
	organizer.UpdateAt = time.Now()

	if err := configs.DB.Save(&user).Error; err != nil {
		return http.StatusInternalServerError, utils.GenerateResponse(false, err.Error())
	}

	if err := configs.DB.Save(&organizer).Error; err != nil {
		return http.StatusInternalServerError, utils.GenerateResponse(false, err.Error())

	}

	return http.StatusOK, utils.GenerateResponse(true, "Success update profile")
}

func UpdatePassword(userID uint, body validators.UpdatePasswordBody) (status int, response gin.H) {
	var user models.MstUser

	if err := configs.DB.Where("id = ? and deleted_at is null", userID).First(&user).Error; err != nil {
		return http.StatusInternalServerError, utils.GenerateResponse(false, err.Error())
	}

	validatePassword := utils.VerifyPassword(user.Password, body.Password)
	if !validatePassword {
		return http.StatusBadRequest, utils.GenerateResponse(false, "Invalid password")
	}

	if body.NewPassword != body.ConfirmNewPassword {
		return http.StatusBadRequest, utils.GenerateResponse(false, "New password and confirm new password must be the same")
	}

	hashedPassword, err := utils.HashPassword(body.NewPassword)

	if err != nil {
		return http.StatusInternalServerError, utils.GenerateResponse(false, err.Error())
	}

	user.Password = hashedPassword

	if err := configs.DB.Save(&user).Error; err != nil {
		return http.StatusInternalServerError, utils.GenerateResponse(false, err.Error())
	}

	return http.StatusOK, utils.GenerateResponse(true, "Success update password")
}

func GetProfile(userID uint) (status int, response gin.H) {
	var user struct {
		ID    uint   `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
		Role  string `json:"role"`
	}

	if err := configs.DB.Table("mst_user u").Where("id = ?", userID).Select("u.id, u.name, u.email, u.role").First(&user).Error; err != nil {
		return http.StatusInternalServerError, utils.GenerateResponse(false, err.Error())
	}

	return http.StatusOK, utils.GenerateResponse(true, "Success get profile", user)
}

func GetOrganizerProfile(userID uint) (status int, response gin.H) {
	var organizer struct {
		ID          uint   `json:"id"`
		Name        string `json:"name"`
		Email       string `json:"email"`
		Role        string `json:"role"`
		CompanyName string `json:"company_name"`
		ContactInfo string `json:"contact_info"`
	}

	if err := configs.DB.Table("mst_user u").Joins("left join mst_organizer o on u.id = o.user_id").Select("u.id, u.name, u.email, u.role, o.company_name, o.contact_info").Where("u.id = ?", userID).First(&organizer).Error; err != nil {
		return http.StatusInternalServerError, utils.GenerateResponse(false, err.Error())
	}

	return http.StatusOK, utils.GenerateResponse(true, "Success get profile", organizer)
}
