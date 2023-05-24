package db

import (
	"context"

	"github.com/24online24/GroupTaskOrganizer-API/pkg/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx context.Context = context.TODO()

func ConnectToAtlas() *mongo.Client {
	uri := utils.GetEnvVariable("MONGO_URI")
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	utils.HandleError(err)

	database := client.Database("GTO_DB")
	roomsCollection = database.Collection("rooms")
	tasksCollection = database.Collection("tasks")
	return client
}
