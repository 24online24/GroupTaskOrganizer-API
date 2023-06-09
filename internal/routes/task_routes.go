package routes

import (
	"github.com/24online24/GroupTaskOrganizer-API/internal/models"
	"github.com/24online24/GroupTaskOrganizer-API/pkg/db"
	"github.com/gin-gonic/gin"
)

// RegisterTaskRoutes înregistrează rutele pentru task-uri
func RegisterTaskRoutes(router *gin.Engine) {
	router.GET("/task/:roomID", getTasksHandler)
	router.POST("/task", createTaskHandler)
	router.PUT("/task/:taskID", updateTaskHandler)
	router.PATCH("/task/:taskID", completeTaskHandler)
	router.DELETE("/task/:taskID", deleteTaskHandler)
}

// getTasksHandler returnează toate task-urile dintr-o cameră
func getTasksHandler(c *gin.Context) {
	roomID := c.Param("roomID")
	tasks, err := db.ReadAllTasksInRoom(roomID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"tasks": tasks})
}

// createTaskHandler creează un task nou
func createTaskHandler(c *gin.Context) {
	var task models.Task
	err := c.BindJSON(&task)
	if err != nil { // dacă nu s-a putut citi body-ul request-ului
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	result, err := db.CreateTask(&task)
	if err != nil { // dacă nu s-a putut crea task-ul
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"taskAdded": true, "task": result})
}

// updateTaskHandler actualizează un task
func updateTaskHandler(c *gin.Context) {
	taskID := c.Param("taskID")
	var task models.Task
	err := c.BindJSON(&task)
	if err != nil { // dacă nu s-a putut citi body-ul request-ului
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	result, err := db.UpdateTask(taskID, &task)
	if err != nil { // dacă nu s-a putut actualiza task-ul
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"taskUpdated": true, "task": result})
}

// completeTaskHandler marchează un task ca fiind completat
func completeTaskHandler(c *gin.Context) {
	taskID := c.Param("taskID")
	var completed models.Task
	err := c.BindJSON(&completed)
	if err != nil { // dacă nu s-a putut citi body-ul request-ului
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	result, err := db.CompleteTask(taskID, &completed)
	if err != nil { // dacă nu s-a putut actualiza task-ul
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, result)
}

// deleteTaskHandler șterge un task
func deleteTaskHandler(c *gin.Context) {
	taskID := c.Param("taskID")
	deleteResult, err := db.DeleteTask(taskID)
	if err != nil || deleteResult.DeletedCount == 0 { // dacă nu s-a șters niciun task
		c.JSON(500, gin.H{"taskDeleted": false})
		return
	}
	c.JSON(200, gin.H{"taskDeleted": true})
}
