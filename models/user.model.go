package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:128;not null" json:"name"`
	Email     string    `gorm:"size:128;not null" json:"email"`
	Password  string    `gorm:"size:128;not null" json:"password"`
	Role      string    `gorm:"size:32;default:'customer'; not null" json:"role"`
	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt time.Time `gorm:"" json:"deleted_at"`
}
