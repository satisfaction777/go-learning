package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

var users = []User{
{ID: 1, Name: "Вика"},
{ID: 2, Name: "Павел"},
}


func getUsers(c *gin.Context) { // вывести список пользователей
	c.JSON(200,users)
}


func addUser(c *gin.Context) { // добавить пользователя
	var user User
	c.ShouldBindJSON(&user)

	if user.Name == "" {
		c.JSON(400, gin.H{"error": "имя не может быть пустым"})
		return
	}
	user.ID = len(users) + 1
	users = append(users, user) // добавить в слайс
	c.JSON(201, user) // вернуть созданного пользователя
}



func deleteUser(c *gin.Context) { // удалить пользователя
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "неверный id"})
		return
}
	for i, u := range users{
		if u.ID == id {
			users = append(users[:i], users[i+1:]...) // удалить
			c.Status(204) // ответ
			return
		}
	}
	c.JSON(404, gin.H{"error": "не найден"}) // если не нашли пользователя
}


func main() {
	r := gin.Default()

	r.GET("/users", getUsers)
	r.POST("/users", addUser)
	r.DELETE("/users/:id", deleteUser)

	r.Run(":8080") // ЗАПУСК на порту
}