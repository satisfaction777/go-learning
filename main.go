package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "strconv"
    "strings"

	"gorm.io/driver/postgres"
    "gorm.io/gorm"
	
)

type Product struct {
	ID int `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Price float64 `json:"price"`
	Stock int `json:"stock"`
}

var db *gorm.DB

func connectDB(){
	dsn := "host=localhost user=admin password=secret dbname=goapp port=5432 sslmode=disable"
	
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		panic("Не удалось подключиться!" + err.Error())
	}
	db.AutoMigrate(&Product{})
	fmt.Println("База готова!")
}


func usersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		var products []Product
		db.Find(&products)
		json.NewEncoder(w).Encode(products)
	
	
	case http.MethodPost:
		var product Product
		json.NewDecoder(r.Body).Decode(&product)
		if product.Name == "" || product.Price <= 0 || product.Stock <= -1  {
			w.WriteHeader(400)
			return
		}
		db.Create(&product)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(product)
	

	case http.MethodDelete:
		idStr := strings.TrimPrefix(r.URL.Path, "/products/")
		id, err := strconv.Atoi(idStr)
		if err != nil{
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		result := db.Delete(&Product{},id)
		if result.RowsAffected == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusNoContent)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}	
}

func main(){
	connectDB()
	http.HandleFunc("/products", usersHandler)
	http.HandleFunc("/products/", usersHandler)
	fmt.Println("Сервер на :8080")
	http.ListenAndServe(":8080", nil)
}