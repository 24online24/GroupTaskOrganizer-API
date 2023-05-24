package db

import (
	"github.com/24online24/GroupTaskOrganizer-API/internal/models"
	"github.com/24online24/GroupTaskOrganizer-API/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var roomsCollection *mongo.Collection

func ReadAllRooms() ([]*models.Room, error) {
	var rooms []*models.Room
	cursor, err := roomsCollection.Find(ctx, bson.D{})
	utils.HandleError(err)

	for cursor.Next(ctx) {
		var room models.Room
		err := cursor.Decode(&room)
		utils.HandleError(err)
		rooms = append(rooms, &room)
	}
	utils.HandleError(err)
	cursor.Close(ctx)
	return rooms, err
}

func ReadRoomByName(name string) (*models.Room, error) {
	var room models.Room
	filter := bson.D{{Key: "name", Value: name}}
	err := roomsCollection.FindOne(ctx, filter).Decode(&room)
	if err != nil {
		return nil, utils.ErrNoDocument
	}
	return &room, err
}

func CreateRoom(name string) (*models.Room, error) {
	searchedRoom, _ := ReadRoomByName(name)
	if searchedRoom != nil {
		return nil, utils.ErrDuplicateDocument
	}
	room := models.Room{Name: name}
	result, err := roomsCollection.InsertOne(ctx, room)
	utils.HandleError(err)
	room.ID = result.InsertedID.(primitive.ObjectID)
	return &room, err

}
