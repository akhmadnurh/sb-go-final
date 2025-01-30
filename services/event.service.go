package services

import (
	"OneTix/configs"
	"OneTix/models"
	"OneTix/utils"
	"OneTix/validators"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func FindEvents(query validators.FindEventsQuery) (status int, response gin.H) {
	var events []struct {
		ID                uint   `json:"id"`
		Name              string `json:"name"`
		Description       string `json:"description"`
		OrganizerName     string `json:"organizer_name"`
		ContactInfo       string `json:"contact_info"`
		Location          string `json:"location"`
		Quota             uint   `json:"quota"`
		StartTime         string `json:"start_time"`
		EndTime           string `json:"end_time"`
		RegistrationStart string `json:"registration_start"`
		RegistrationEnd   string `json:"registration_end"`
	}

	offset := utils.GetOffset(query.Page, query.Limit)

	if err := configs.DB.Table("mst_event e").Joins("left join mst_organizer o on e.organizer_id = o.id").Select("e.id, e.name, e.description, o.company_name as organizer_name, o.contact_info, e.location, e.quota, e.start_time, e.end_time, e.registration_start, e.registration_end").Where("e.deleted_at is null").Offset(offset).Limit(query.Limit).Find(&events).Error; err != nil {
		return http.StatusInternalServerError, utils.GenerateResponse(false, err.Error())
	}

	return http.StatusOK, utils.GenerateResponse(true, "Success", events)
}

func CreateEvent(userID uint, body validators.CreateEventBody) (status int, response gin.H) {
	var organizer models.MstOrganizer

	if err := configs.DB.Where("user_id = ? and deleted_at is null", userID).First(&organizer).Error; err != nil {
		fmt.Println(err)
		return http.StatusInternalServerError, utils.GenerateResponse(false, err.Error())
	}

	newEvent := models.MstEvent{
		Name:              body.Name,
		Description:       body.Description,
		OrganizerID:       organizer.ID,
		Location:          body.Location,
		Quota:             body.Quota,
		StartTime:         body.StartTime,
		EndTime:           body.EndTime,
		RegistrationStart: body.RegistrationStart,
		RegistrationEnd:   body.RegistrationEnd,
	}

	if err := configs.DB.Create(&newEvent).Error; err != nil {
		return http.StatusInternalServerError, utils.GenerateResponse(false, err.Error())
	}

	return http.StatusOK, utils.GenerateResponse(true, "Success create event")
}

func DeleteEvent(query validators.DeleteEventQuery) (status int, response gin.H) {
	var event models.MstEvent

	if err := configs.DB.Where("id = ? and deleted_at is null", query.ID).First(&event).Error; err != nil {
		if err.Error() == "record not found" {
			return http.StatusNotFound, utils.GenerateResponse(false, "Event not found")
		}

		return http.StatusInternalServerError, utils.GenerateResponse(false, err.Error())
	}

	if err := configs.DB.Model(event).Update("deleted_at", time.Now()).Error; err != nil {
		return http.StatusInternalServerError, utils.GenerateResponse(false, err.Error())
	}

	return http.StatusOK, utils.GenerateResponse(true, "Success delete event")
}

func UpdateEvent(query validators.UpdateEventQuery, body validators.UpdateEventBody) (status int, response gin.H) {
	var event models.MstEvent

	if err := configs.DB.Where("id = ? and deleted_at is null", query.ID).First(&event).Error; err != nil {
		if err.Error() == "record not found" {
			return http.StatusNotFound, utils.GenerateResponse(false, "Event not found")
		}

		return http.StatusInternalServerError, utils.GenerateResponse(false, err.Error())
	}

	if err := configs.DB.Model(event).Updates(models.MstEvent{Name: body.Name, Description: body.Description, Location: body.Location, Quota: body.Quota, StartTime: body.StartTime, EndTime: body.EndTime, RegistrationStart: body.RegistrationStart, RegistrationEnd: body.RegistrationEnd, UpdatedAt: time.Now()}).Error; err != nil {
		return http.StatusInternalServerError, utils.GenerateResponse(false, err.Error())

	}

	return http.StatusOK, utils.GenerateResponse(true, "Success update event")
}
