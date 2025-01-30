package models

import (
	"time"

	"gorm.io/gorm"
)

type MstUser struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"size:128;not null" json:"name"`
	Email     string         `gorm:"size:128;not null" json:"email"`
	Password  string         `gorm:"size:128;not null" json:"password"`
	Role      string         `gorm:"size:32;default:'customer'; not null" json:"role"`
	CreatedAt time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	MstOrganizer MstOrganizer `gorm:"foreignKey:UserID"`
	TrsTicket    []TrsTicket  `gorm:"foreignKey:UserID"`
}

func (MstUser) TableName() string {
	return "mst_user"
}
