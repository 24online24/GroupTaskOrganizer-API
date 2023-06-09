package db

import (
	"context"

	"github.com/24online24/GroupTaskOrganizer-API/pkg/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx context.Context = context.TODO()

func ConnectToAtlas() *mongo.Client {
	uri := utils.GetEnvVariable("MONGO_URI")                                       // obține URI-ul de conectare la Atlas
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)                      // setează versiunea API-ului de server
	clientOptions := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI) // setează opțiunile clientului

	client, err := mongo.Connect(context.TODO(), clientOptions) // conectează clientul la baza de date
	utils.HandleError(err)                                      // dacă apare o eroare, o afișează

	database := client.Database("GTO_DB")          // obține baza de date
	roomsCollection = database.Collection("rooms") // obține colecția de "camere"
	tasksCollection = database.Collection("tasks") // obține colecția de task-uri
	return client
}
