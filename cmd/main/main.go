package main

import (
	"context"
	"log"

	"github.com/24online24/GroupTaskOrganizer-API/internal/routes"
	"github.com/24online24/GroupTaskOrganizer-API/pkg/db"
	"github.com/24online24/GroupTaskOrganizer-API/pkg/utils"
)

func main() {
	log.Default().Println("Starting server on port 8080")
	mongoClient := db.ConnectToAtlas()
	defer func() {
		utils.HandleError(mongoClient.Disconnect(context.TODO()))
	}()
	log.Default().Println("Connected to MongoDB Atlas")
	router := routes.NewRouter()
	routes.RegisterRoomRoutes(router)
	routes.RegisterTaskRoutes(router)
	err := router.Run(":8000")
	utils.HandleError(err)
}
