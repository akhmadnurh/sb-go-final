package models

import "time"

type TrsTicket struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	EventID      uint      `gorm:"not null" json:"event_id"`
	UserID       uint      `gorm:"not null" json:"user_id"`
	TicketCode   string    `gorm:"size:32;not null" json:"ticket_code"`
	IsUsed       bool      `gorm:"not null;default:false" json:"is_used"`
	RegisteredAt time.Time `gorm:"not null" json:"registered_at"`
	CreatedAt    time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt    time.Time `gorm:"index" json:"deleted_at,omitempty"`
}

func (TrsTicket) TableName() string {
	return "trs_ticket"

}
