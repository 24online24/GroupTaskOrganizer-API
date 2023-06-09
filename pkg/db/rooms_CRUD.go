package db

import (
	"github.com/24online24/GroupTaskOrganizer-API/internal/models"
	"github.com/24online24/GroupTaskOrganizer-API/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var roomsCollection *mongo.Collection

// ReadAllRooms obține toate camerele din colecție
func ReadAllRooms() ([]*models.Room, error) {
	var rooms []*models.Room                           // slice de pointeri la structuri de tipul Room, inițial gol
	cursor, err := roomsCollection.Find(ctx, bson.D{}) // obține toate documentele din colecție
	utils.HandleError(err)

	for cursor.Next(ctx) { // pentru fiecare document
		var room models.Room        // inițializează o structură de tipul Room
		err := cursor.Decode(&room) // decodifică documentul
		utils.HandleError(err)
		rooms = append(rooms, &room) // adaugă structura la slice-ul definit la începutul funcției
	}
	utils.HandleError(err)
	cursor.Close(ctx)
	return rooms, err
}

// ReadRoomByName obține o cameră după nume
func ReadRoomByName(name string) (*models.Room, error) {
	var room models.Room
	filter := bson.D{{Key: "name", Value: name}}              // filtrează după nume
	err := roomsCollection.FindOne(ctx, filter).Decode(&room) // obține un singur document
	if err != nil {
		return nil, utils.ErrNoDocument
	}
	return &room, err
}

// CreateRoom creează o cameră nouă
func CreateRoom(name string) (*models.Room, error) {
	searchedRoom, _ := ReadRoomByName(name) // caută o cameră cu același nume
	if searchedRoom != nil {                // dacă există deja o cameră cu același nume
		return nil, utils.ErrDuplicateDocument // returnează eroare
	}
	room := models.Room{Name: name}                     // inițializează o structură de tipul Room
	result, err := roomsCollection.InsertOne(ctx, room) // inserează documentul în colecție
	utils.HandleError(err)
	room.ID = result.InsertedID.(primitive.ObjectID) // obține ID-ul documentului introdus
	return &room, err

}
