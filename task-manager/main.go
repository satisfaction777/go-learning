package main

import (
"task-manager/db"        
"task-manager/handlers"
"github.com/gin-gonic/gin"
)



func main() {
db.ConnectDB()

r := gin.Default()

r.GET("/tasks/:id", handlers.GetTask)
r.GET("/tasks", handlers.GetTasks)
r.POST("/tasks", handlers.CreateTask)
r.PUT("/tasks/:id", handlers.UpdateTask)
r.DELETE("/tasks/:id", handlers.DeleteTask)

r.Run(":8080")

}