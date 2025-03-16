package models

type User struct {
	ID       uint   `gorm:"primary_key"`
	Name     string `gorm:"unique"`
	Email    string `gorm:"unique"`
	Password string `gorm:"not null"`
}
