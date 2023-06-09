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
	defer func() { // închide conexiunea la baza de date la terminarea programului
		utils.HandleError(mongoClient.Disconnect(context.TODO()))
	}()
	log.Default().Println("Connected to MongoDB Atlas")
	router := routes.NewRouter()      // crează un nou router
	routes.RegisterRoomRoutes(router) // înregistrează rutele pentru "camere"
	routes.RegisterTaskRoutes(router) // înregistrează rutele pentru task-uri
	err := router.Run(":8000")        // pornește server-ul pe portul 8000
	utils.HandleError(err)            // dacă apare o eroare, o afișează
}
