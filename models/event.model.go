package models

import (
	"time"

	"gorm.io/gorm"
)

type MstEvent struct {
	ID                uint           `gorm:"primaryKey" json:"id"`
	Name              string         `gorm:"size:192;not null" json:"name"`
	Description       string         `gorm:"size:256;not null;type:text" json:"description"`
	OrganizerID       uint           `gorm:"not null" json:"organizer_id"`
	Location          string         `gorm:"size:128;not null" json:"location"`
	Quota             uint           `gorm:"not null" json:"quota"`
	StartTime         time.Time      `gorm:"not null" json:"start_time"`
	EndTime           time.Time      `gorm:"not null" json:"end_time"`
	RegistrationStart time.Time      `gorm:"not null" json:"registration_start"`
	RegistrationEnd   time.Time      `gorm:"not null" json:"registration_end"`
	CreatedAt         time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt         time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	TrsTicket []TrsTicket `gorm:"foreignKey:EventID"`
}

func (MstEvent) TableName() string {
	return "mst_event"
}
