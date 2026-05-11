package main

import (
"task-manager/db"        
"task-manager/handlers"
"github.com/gin-gonic/gin"
"fmt"
)


func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Printf("→ %s %s\n", c.Request.Method, c.Request.URL.Path)
		c.Next()
		fmt.Printf("← %d\n", c.Writer.Status())
	}
}


func main() {
db.ConnectDB()

r := gin.Default()

r.Use(Logger())

r.GET("/tasks/:id", handlers.GetTask)
r.GET("/tasks", handlers.GetTasks)
r.POST("/tasks", handlers.CreateTask)
r.PUT("/tasks/:id", handlers.UpdateTask)
r.DELETE("/tasks/:id", handlers.DeleteTask)

r.Run(":8080")

}