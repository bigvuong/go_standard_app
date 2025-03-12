package models

import (
	"time"
	"standard_app/config"
)

type Product struct {
	ID   		uint   `json:"id" gorm:"primaryKey"`
	Name 		string `json:"name" gorm:"unique;not null"`
	Quantity 	int `json:"quantity"`
	Price 		float64 `json:"price"`
	Slug 		string `json:"slug"`
	UserID 		uint   `json:"user_id" gorm:"not null"`
    User   		*User `json:"user" gorm:"foreignKey:UserID;references:ID"`
	CreatedAt 	time.Time `json:"created_at"`
	CreatedBy 	uint `json:"created_by"`
	UpdatedAt 	time.Time `json:"updated_at"`
	UpdatedBy 	uint `json:"updated_by"`
	DeletedAt 	*time.Time `json:"deleted_at"`
	DeletedBy 	uint `json:"deleted_by"`
}

func (p *Product) Save() (*Product, error) {
	if err := config.DB.Create(p).Error; err != nil {
		return nil, err
	}
	return p, nil
}

func (p *Product) GetAll() ([]Product, error) {
	var products []Product
	if err := config.DB.Preload("User").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (p *Product) GetById(id string) (*Product, error) {
	var product *Product
	if err := config.DB.Preload("User").First(&product, id).Error; err != nil {
		return nil, err
	}
	return product, nil
}
