package db


import (
	"gorm.io/driver/postgres"
    "gorm.io/gorm"
	"task-manager/models"
	"fmt"
)


var db *gorm.DB



func ConnectDB() {
	dsn := "host=localhost user=admin password=secret dbname=goapp port=5432 sslmode=disable"

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