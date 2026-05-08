package handlers


import (
	"github.com/gin-gonic/gin"
	"task-manager/db"
	"task-manager/models"
)



func GetTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	result := db.GetDB().First(&task, id)
	if result.Error != nil {
		c.JSON(404,gin.H{"error": "задача не найдена"})
		return
	}
	c.JSON(200, task)
}



func GetTasks(c *gin.Context) {
	var tasks []models.Task
	db.GetDB().Find(&tasks)
	c.JSON(200, tasks)
}



func CreateTask(c *gin.Context) {
	var task models.Task
	c.ShouldBindJSON(&task)

	if task.Title == "" {
		c.JSON(400, gin.H{"error": "название задачи не может быть пустым"})
		return
	}
	db.GetDB().Create(&task)
	c.JSON(201, task)
}



func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	result := db.GetDB().First(&task, id)
	
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "задача не найдена"})
		return
	}
	var newTask models.Task
	c.ShouldBindJSON(&newTask)
	if newTask.Title == "" {
		c.JSON(400, gin.H{"error": "название задачи не может быть пустым"})
		return
	}
	db.GetDB().Model(&task).Updates(newTask)
	c.JSON(200, newTask) 
}



func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	result := db.GetDB().Delete(&task, id)
	
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "ошибка при удалении"})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{"error": "такой задачи нет"})
		return
	}
	c.Status(204)
}