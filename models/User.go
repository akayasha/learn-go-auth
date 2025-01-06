package models

type User struct {
	ID              uint   `gorm:"primaryKey"`
	Username        string `gorm:"size:255;not null"`
	Email           string `gorm:"type:varchar(255);uniqueIndex;not null"`
	PasswordHash    string `gorm:"not null"`
	Role            string `gorm:"size:50;default:'user'"`
	OTP             string `gorm:"size:6"`
	IsEmailVerified bool   `gorm:"default:false"`
}
