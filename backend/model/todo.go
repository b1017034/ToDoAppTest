package model

import "time"

type ToDo struct {
	ID        uint `gorm:"primaryKey"`
	Text      string
	UpdatedAT time.Time
	CreatedAT time.Time
}
