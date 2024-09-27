package models
import "time"

type User struct {
	ID        int       `json:"id" gorm:"primary_key"`
	FullName  string    `json:"full_name" gorm:"not null "`
	Username  string    `json:"username" gorm:"unique"`
	Password  string    `json:"password" gorm:"not null"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	IsBlocked bool      `json:"is_blocked" gorm:"default:false"`
	IsDeleted bool      `json:"is_deleted" gorm:"default:false"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type SwagUser struct {
	FullName string `json:"full_name"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password" gorm:"not null"`
}


type UserFilterForUpdate struct {
	FullName  string    `json:"full_name" gorm:"not null "`
	Username  string    `json:"username" gorm:"unique"`
	
	
}

