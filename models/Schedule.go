package models

import "time"

type Schedule struct {
	ScheduleId uint32 `gorm:"primaryKey" json:"id,omitempty"`
	UserID uint32 `json:"user_id,omitempty" gorm:"AUTO_INCREMENT"`
	Title string `json:"title,omitempty" gorm:"type:varchar(64)"`
	Day string `json:"day,omitempty" gorm:"type:varchar(9)"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}