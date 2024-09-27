package models

import "time"

type Announcement struct {
	ID                int       `json:"id" gorm:"primary_key"`
	Price             int       `json:"price" gorm:"not null"`
	CountApart        int       `json:"count_apart" gorm:"not null"`
	Floor             int       `json:"floor" gorm:"not null"`
	SquareMeters      int       `json:"square_meters" gorm:"not null"`
	District          string    `json:"district" gorm:"not null"`
	Pets              string    `json:"pets" gorm:"not null"`
	TypeOfDevelopment string    `json:"typeofdevelopment" gorm:"not null"`
	Comment           string    `json:"comment" `
	Phone             string    `json:"phone" `
	CreatedBy         int       `json:"created_by" gorm:"not null"`
	CreatedAt         time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt         time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	ViewCount         int64       `json:"view_count" gorm:"view_count"`
}

type AnnouncementFilterForUpdate struct {
	Price             int    `json:"price" gorm:"not null"`
	CountApart        int    `json:"count_apart" gorm:"not null"`
	Floor             int    `json:"floor" gorm:"not null"`
	SquareMeters      int    `json:"square_meters" gorm:"not null"`
	District          string `json:"district" gorm:"not null"`
	Pets              string `json:"pets" gorm:"not null"`
	TypeOfDevelopment string `json:"typeofdevelopment" gorm:"not null"`
	Comment           string `json:"comment" `
	Phone             string `json:"phone" `
}

type AnnouncementView struct {
	ID             uint `json:"id" gorm:"primary_key"`
	AnnouncementID uint `json:"announcement_id"`
	UserID         uint `json:"user_id"`
}
