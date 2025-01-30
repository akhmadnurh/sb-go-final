package services

import (
	"OneTix/configs"
	"OneTix/models"
	"OneTix/utils"
	"OneTix/validators"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func FindTickets(userID uint, query validators.FindTicketsQuery) (status int, response gin.H) {
	var tickets []struct {
		ID               uint   `json:"id"`
		EventID          uint   `json:"event_id"`
		TicketCode       string `json:"ticket_code"`
		IsUsed           bool   `json:"is_used"`
		RegisteredAt     string `json:"registered_at"`
		EventName        string `json:"event_name"`
		EventDescription string `json:"event_description"`
		EventLocation    string `json:"event_location"`
		StartTime        string `json:"start_time"`
		EndTime          string `json:"end_time"`
	}

	offset := utils.GetOffset(query.Page, query.Limit)

	if err := configs.DB.Table("trs_ticket t").Joins("left join mst_event e on t.event_id = e.id").Select("t.id, t.event_id, t.ticket_code, t.is_used, t.registered_at, e.name as event_name, e.description as event_description, e.location as event_location, e.start_time, e.end_time").Where("t.user_id = ? and t.deleted_at is null and e.deleted_at is null", userID).Offset(offset).Limit(query.Limit).Find(&tickets).Error; err != nil {
		return http.StatusInternalServerError, utils.GenerateResponse(false, err.Error())
	}

	return http.StatusOK, utils.GenerateResponse(true, "Success get tickets", tickets)
}

func CreateTicket(userID uint, body validators.CreateTicketBody) (status int, response gin.H) {
	var event models.MstEvent
	var ticketCount int64

	now := time.Now()

	if err := configs.DB.Where("id = ? and deleted_at is null", body.EventID).First(&event).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return http.StatusNotFound, utils.GenerateResponse(false, "Event not found")
		}

		return http.StatusInternalServerError, utils.GenerateResponse(false, err.Error())
	}

	if err := configs.DB.Model(models.TrsTicket{}).Where("event_id = ? and deleted_at is null", body.EventID).Count(&ticketCount).Error; err != nil {
		return http.StatusInternalServerError, utils.GenerateResponse(false, err.Error())
	}

	if now.Before(event.RegistrationStart) || now.After(event.RegistrationEnd) {
		return http.StatusBadRequest, utils.GenerateResponse(false, "Registration is not open or closed")
	}

	if uint(ticketCount) > event.Quota {
		return http.StatusBadRequest, utils.GenerateResponse(false, "Quota is full")
	}

	newTicket := models.TrsTicket{
		EventID:      body.EventID,
		UserID:       userID,
		TicketCode:   utils.GenerateTicketCode(),
		RegisteredAt: now,
		IsUsed:       false,
	}

	if err := configs.DB.Create(&newTicket).Error; err != nil {
		return http.StatusInternalServerError, utils.GenerateResponse(false, err.Error())
	}

	return http.StatusOK, utils.GenerateResponse(true, "Success create ticket")
}

func CheckIn(userID uint, body validators.CheckInBody) (status int, response gin.H) {
	var ticket models.TrsTicket
	var event models.MstEvent

	now := time.Now()

	if err := configs.DB.Where("event_id = ? and user_id = ? and ticket_code = ? and deleted_at is null", body.EventID, userID, body.TicketCode).First(&ticket).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return http.StatusNotFound, utils.GenerateResponse(false, "Ticket not found")
		}

		return http.StatusInternalServerError, utils.GenerateResponse(false, err.Error())
	}

	if err := configs.DB.Where("id = ? and deleted_at is null", body.EventID).First(&event).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return http.StatusNotFound, utils.GenerateResponse(false, "Event not found")
		}

		return http.StatusInternalServerError, utils.GenerateResponse(false, err.Error())
	}

	if now.Before(event.StartTime) || now.After(event.EndTime) {
		return http.StatusBadRequest, utils.GenerateResponse(false, "The event is not open yet or has already closed.")
	}

	ticket.IsUsed = true
	ticket.UpdatedAt = now

	if err := configs.DB.Save(&ticket).Error; err != nil {
		return http.StatusInternalServerError, utils.GenerateResponse(false, err.Error())

	}

	return http.StatusOK, utils.GenerateResponse(true, "Success check in")
}
