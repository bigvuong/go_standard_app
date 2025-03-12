package databases

import (
    "fmt"
    "gorm.io/gorm"
	"standard_app/models"
)


func SeedData(db *gorm.DB) {
	// Seed Users
	users := []models.User{
		{UserName: "johndoe", Password: "12345678", Name: "John Doe", Email: "john.doe@example.com", Phone: "12222222", Age: 20},
		{UserName: "janesmith", Password: "12345678", Name: "Jane Smith", Email: "jane.smith@example.com", Phone: "12222223", Age: 34},
		{UserName: "alicejohnson", Password: "12345678", Name: "Alice Johnson", Email: "alice.johnson@example.com", Phone: "122222423", Age: 25},
	}

	// Check if users exist, if not, insert them
	for _, user := range users {
		if err := db.Where("user_name = ?", user.UserName).First(&user).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				if err := db.Create(&user).Error; err != nil {
					fmt.Println("Error seeding user data: %v", err)
				}
				fmt.Println("User %s created", user.Name)
			} else {
				fmt.Println("Error checking if user exists: %v", err)
			}
		}
	}

	// Seed Products
	products := []models.Product{
		{Name: "Product A", Quantity: 100, Price: 19.99, Slug: "product-a", UserID: 1},
		{Name: "Product B", Quantity: 300, Price: 49.99, Slug: "product-b", UserID: 1},
		{Name: "Product C", Quantity: 150, Price: 29.99, Slug: "product-c", UserID: 2},
	}

	// Check if products exist, if not, insert them
	for _, product := range products {
		if err := db.Where("name = ?", product.Name).First(&product).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				if err := db.Create(&product).Error; err != nil {
					fmt.Println("Error seeding product data: %v", err)
				}
				fmt.Println("Product %s created", product.Name)
			} else {
				fmt.Println("Error checking if product exists: %v", err)
			}
		}
	}
}