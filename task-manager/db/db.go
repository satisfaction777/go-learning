package db


import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
    "gorm.io/gorm"
	"task-manager/models"
	"fmt"
)


var db *gorm.DB



func ConnectDB() {

	godotenv.Load()

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")


	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	
	if err != nil {
		panic("не удалось подключиться: " + err.Error())
	}

	db.AutoMigrate(&models.Task{})
	fmt.Println("База готова!")
}



func GetDB() *gorm.DB {
	return db
}