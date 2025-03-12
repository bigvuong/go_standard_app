package databases

import (
    "fmt"
    "log"
    "gorm.io/gorm"
	"standard_app/config"
	"standard_app/models"
)

var gTables = []string{
    "products",
    "users", // Specify table names in a slice
}

var gMigrateModels = map[string]interface{}{
    "users": &models.User{}, // List of models as interfaces
    "products": &models.Product{},
}

// Initialize and run migrations
func RunMigrations() {
    // Connect to the database
    config.LoadConfig()

    // Drop tables before AutoMigrate
    fmt.Println("======= Tables Dropped start ...")
    err := DropTables(config.DB)
    if err != nil {
        log.Fatal("Failed to drop tables:", err)
    }
    fmt.Println("======= Tables Dropped successfully!")

    fmt.Println("======= Database Migration start...")
    // Run migrations for the models
    for table, gModel := range gMigrateModels {
        fmt.Println("======= Migration Table: ", table)
        config.DB.AutoMigrate(gModel)
    }
    fmt.Println("======= Database Migration completed successfully!")


    fmt.Println("======= Database Seed start...")
    SeedData(config.DB)
    fmt.Println("======= Database Seed completed successfully!")
}

func DropTables(db *gorm.DB) error {
    // Start a transaction
    tx := db.Begin()

    // Drop the tables (you can also dynamically fetch the table names if needed)

    for _, gTable := range gTables {
        fmt.Println("======= Drop Table: %s", gTable)
        if err := tx.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s;", gTable)).Error; err != nil {
            tx.Rollback()
            return fmt.Errorf("failed to drop table %s: %v", gTable, err)
        }
    }

    // Commit the transaction
    return tx.Commit().Error
}