package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"hackinroms.com/books/models"
	"log"
	"os"
)

var DB *gorm.DB

// Setup db configs production
func Setup() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_DATABASE")
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")

	db, err := gorm.Open("postgres",
		"host="+host+
			" port="+port+
			" user="+user+
			" dbname="+dbname+
			" sslmode=disable password="+password)

	if err != nil {
		log.Fatal(err)
	}

	db.LogMode(false)

	db.AutoMigrate([]models.Book{})
	db.AutoMigrate([]models.User{})

	DB = db
}

// SetupTest test environment, can have different db test environments
func SetupTest() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_DATABASE_TEST")
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")

	db, err := gorm.Open("postgres",
		"host="+host+
			" port="+port+
			" user="+user+
			" dbname="+dbname+
			" sslmode=disable password="+password)

	if err != nil {
		log.Fatal(err)
	}

	db.LogMode(false)

	db.AutoMigrate([]models.Book{})
	db.AutoMigrate([]models.User{})

	DB = db
}

// GetDB return DB connection
func GetDB() *gorm.DB {
	return DB
}
