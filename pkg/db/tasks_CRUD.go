package db

import (
	"github.com/24online24/GroupTaskOrganizer-API/internal/models"
	"github.com/24online24/GroupTaskOrganizer-API/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var tasksCollection *mongo.Collection

// ReadAllTasksInRoom returnează toate task-urile dintr-o cameră
func ReadAllTasksInRoom(roomID string) ([]*models.Task, error) {
	var tasks []*models.Task
	cursor, err := tasksCollection.Find(ctx, bson.D{})
	utils.HandleError(err)

	for cursor.Next(ctx) { // pentru fiecare document
		var task models.Task
		err := cursor.Decode(&task) // decodifică documentul
		utils.HandleError(err)
		tasks = append(tasks, &task) // adaugă structura la slice-ul definit la începutul funcției
	}
	utils.HandleError(err)
	cursor.Close(ctx)
	return tasks, err
}

// CreateTask creează un task nou
func CreateTask(task *models.Task) (*models.Task, error) {
	result, err := tasksCollection.InsertOne(ctx, task) // inserează documentul în colecție
	utils.HandleError(err)
	task.ID = result.InsertedID.(primitive.ObjectID)
	return task, err // returnează documentul inserat
}

// UpdateTask actualizează un task
func UpdateTask(taskID string, task *models.Task) (*models.Task, error) {
	id, err := primitive.ObjectIDFromHex(taskID)
	utils.HandleError(err)
	filter := bson.D{{Key: "_id", Value: id}} // filtrează după ID
	update := bson.D{
		{Key: "$set",
			Value: bson.D{
				{Key: "title", Value: task.Title},
				{Key: "description", Value: task.Description},
			}}} // actualizează titlul și descrierea
	_, err = tasksCollection.UpdateOne(ctx, filter, update) // actualizează documentul
	utils.HandleError(err)
	task.ID = id
	return task, err // returnează documentul actualizat
}

// CompleteTask marchează un task ca fiind completat
func CompleteTask(taskID string, task *models.Task) (*models.Task, error) {
	id, err := primitive.ObjectIDFromHex(taskID) // obține ID-ul task-ului
	utils.HandleError(err)
	filter1 := bson.D{{Key: "_id", Value: id}} // filtrează după ID
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "completed", Value: task.Completed},
		}}} // actualizează câmpul "completed"
	_, err = tasksCollection.UpdateOne(ctx, filter1, update) // actualizează documentul
	utils.HandleError(err)
	var taskToReturn models.Task
	filter2 := bson.D{{Key: "_id", Value: id}}                        // filtrează după ID
	err = tasksCollection.FindOne(ctx, filter2).Decode(&taskToReturn) // obține documentul actualizat
	utils.HandleError(err)
	return &taskToReturn, err // returnează documentul actualizat
}

// DeleteTask șterge un task
func DeleteTask(taskID string) (*mongo.DeleteResult, error) {
	id, err := primitive.ObjectIDFromHex(taskID) // obține ID-ul task-ului
	utils.HandleError(err)
	filter := bson.D{{Key: "_id", Value: primitive.ObjectID(id)}} // filtrează după ID
	result, err := tasksCollection.DeleteOne(ctx, filter)         // șterge documentul
	utils.HandleError(err)
	return result, err // returnează rezultatul ștergerii
}
