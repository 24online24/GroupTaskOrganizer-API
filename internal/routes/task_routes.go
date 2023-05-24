package routes

import (
	"github.com/24online24/GroupTaskOrganizer-API/internal/models"
	"github.com/24online24/GroupTaskOrganizer-API/pkg/db"
	"github.com/gin-gonic/gin"
)

func RegisterTaskRoutes(router *gin.Engine) {
	router.GET("/task/:roomID", getTasksHandler)
	router.POST("/task", createTaskHandler)
	router.PUT("/task/:taskID", updateTaskHandler)
	router.PATCH("/task/:taskID", completeTaskHandler)
	router.DELETE("/task/:taskID", deleteTaskHandler)
}

func getTasksHandler(c *gin.Context) {
	roomID := c.Param("roomID")
	tasks, err := db.ReadAllTasksInRoom(roomID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"tasks": tasks})
}

func createTaskHandler(c *gin.Context) {
	var task models.Task
	err := c.BindJSON(&task)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	result, err := db.CreateTask(&task)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"taskAdded": true, "task": result})
}

func updateTaskHandler(c *gin.Context) {
	taskID := c.Param("taskID")
	var task models.Task
	err := c.BindJSON(&task)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	result, err := db.UpdateTask(taskID, &task)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"taskUpdated": true, "task": result})
}

func completeTaskHandler(c *gin.Context) {
	taskID := c.Param("taskID")
	var completed models.Task
	err := c.BindJSON(&completed)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	result, err := db.CompleteTask(taskID, &completed)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, result)
}

func deleteTaskHandler(c *gin.Context) {
	taskID := c.Param("taskID")
	deleteResult, err := db.DeleteTask(taskID)
	if err != nil || deleteResult.DeletedCount == 0 {
		c.JSON(500, gin.H{"taskDeleted": false})
		return
	}
	c.JSON(200, gin.H{"taskDeleted": true})
}
