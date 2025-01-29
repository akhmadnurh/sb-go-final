package models

import "time"

type MstOrganizer struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	UserID      uint      `gorm:"not null" json:"user_id"`
	CompanyName string    `gorm:"size:128;not null" json:"company_name"`
	ContactInfo string    `gorm:"size:16;not null" json:"contact_info"`
	CreatedAt   time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdateAt    time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt   time.Time `gorm:"index" json:"deleted_at,omitempty"`
}

func (MstOrganizer) TableName() string {
	return "mst_organizer"
}
