package models

import (
	"time"
	"standard_app/config"
)

type User struct {
	ID        	uint      `json:"id" gorm:"primaryKey"`
	UserName  	string    `json:"user_name" gorm:"unique;not null"`
	Password  	string    `json:"password"`
	Email     	string    `json:"email" gorm:"not null"`
	Phone     	string    `json:"phone" gorm:"not null"`
	Name     	string    `json:"name"`
	Age     	int    	  `json:"age"`
	CreatedAt 	time.Time `json:"created_at"`
	CreatedBy 	uint `json:"created_by"`
	UpdatedAt 	time.Time `json:"updated_at"`
	UpdatedBy 	uint `json:"updated_by"`
	DeletedAt 	*time.Time `json:"deleted_at"`
	DeletedBy 	uint `json:"deleted_by"`
}

func (u *User) Save() (*User, error) {
	if err := config.DB.Create(u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

func (u *User) GetAll() ([]User, error) {
	var users []User
	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
