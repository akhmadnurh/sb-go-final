package validators

import "time"

type FindEventsQuery struct {
	PaginationQuery
}

type CreateEventBody struct {
	Name              string    `json:"name" validate:"required,min=3,max=128"`
	Description       string    `json:"description" validate:"required,min=3,max=256"`
	Location          string    `json:"location" validate:"required,min=3,max=128"`
	Quota             uint      `json:"quota" validate:"required,min=1"`
	StartTime         time.Time `json:"start_time" validate:"required"`
	EndTime           time.Time `json:"end_time" validate:"required"`
	RegistrationStart time.Time `json:"registration_start" validate:"required"`
	RegistrationEnd   time.Time `json:"registration_end" validate:"required"`
}

type DeleteEventQuery struct {
	ID uint `form:"id" validate:"required"`
}

type UpdateEventBody struct {
	CreateEventBody
}

type UpdateEventQuery struct {
	DeleteEventQuery
}
