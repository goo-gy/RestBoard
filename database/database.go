package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"github.com/restBoard/api/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var err = godotenv.Load(".env")
var dsn = fmt.Sprintf("host=%s user=%s password=%s database.Dbname=%s port=%s sslmode=disable TimeZone=Asia/Seoul",
	os.Getenv("DB_HOST"),
	os.Getenv("DB_USER"),
	os.Getenv("DB_PASSWORD"),
	os.Getenv("DB_NAME"),
	os.Getenv("DB_PORT"),
)
var Db, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{})

func init() {
	Db.AutoMigrate(&domain.Posting{})
	fmt.Println("database.Db Migrated")
}

