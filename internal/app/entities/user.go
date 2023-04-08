package entities

import "time"

type User struct {
	Id        int64     `gorm:"primaryKey" json:"id"`
	Email     string    `gorm:"unique" json:"email"`
	Password  string    `gorm:"not null" json:"password"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
