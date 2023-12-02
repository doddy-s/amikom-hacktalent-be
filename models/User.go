package models

import "time"

type User struct {
	UserId    uint32   `gorm:"primaryKey" gorm:"AUTO_INCREMENT" json:"id"`
	Email     string `json:"email"`
	Schedules []Schedule `json:"-"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
}