package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
    "gorm.io/gorm"
    "gorm.io/driver/mysql"
	"gorm.io/gorm/logger"
)

var (
	DB *gorm.DB
)

func LoadConfig() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to MySQL
	dsn := os.Getenv("DB_DSN") // MySQL DSN in .env (user:password@tcp(localhost:3306)/dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
        DisableForeignKeyConstraintWhenMigrating: false, // Bật foreign key constraint
		Logger: logger.Default.LogMode(logger.Silent), // Tắt tất cả log SQL
    })
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	DB = db
}
