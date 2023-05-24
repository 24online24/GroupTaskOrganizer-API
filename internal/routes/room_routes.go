package routes

import (
	"github.com/24online24/GroupTaskOrganizer-API/internal/models"
	"github.com/24online24/GroupTaskOrganizer-API/pkg/db"
	"github.com/gin-gonic/gin"
)

func RegisterRoomRoutes(router *gin.Engine) {
	router.GET("/rooms", getRoomsHandler)
	router.GET("/rooms/:name", getRoomByNameHandler)
	router.POST("/rooms", createRoomHandler)
}

func getRoomsHandler(c *gin.Context) {
	rooms, err := db.ReadAllRooms()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"rooms": rooms})
}

func getRoomByNameHandler(c *gin.Context) {
	name := c.Param("name")
	room, err := db.ReadRoomByName(name)
	if err != nil {
		c.JSON(500, gin.H{"status": false, "message": "Room not found"})
		return
	}
	c.JSON(200, gin.H{"status": true, "room": room})
}

func createRoomHandler(c *gin.Context) {
	var room *models.Room
	err := c.BindJSON(&room)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	result, err := db.CreateRoom(room.Name)
	if err != nil {
		c.JSON(500, gin.H{"status": false, "message": "Room already exists"})
		return
	}
	c.JSON(200, gin.H{"status": true, "room": result})
}
