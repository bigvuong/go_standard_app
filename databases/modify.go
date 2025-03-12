package databases

import (
    "fmt"
	"standard_app/config"
	"standard_app/models"
)

var gModifyModels = []interface{}{
    &models.User{}, // List of models as interfaces
    &models.Product{},
}

// Initialize and run migrations
func RunModifies() {
    // Connect to the database
    config.LoadConfig()

    fmt.Println("======= Database modify start...")
    // Run migrations for the models
    for _, gModel := range gModifyModels {
        config.DB.AutoMigrate(gModel)
    }
    fmt.Println("======= Database modify completed successfully!")
}