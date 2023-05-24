package db

import (
	"github.com/24online24/GroupTaskOrganizer-API/internal/models"
	"github.com/24online24/GroupTaskOrganizer-API/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var tasksCollection *mongo.Collection

func ReadAllTasksInRoom(roomID string) ([]*models.Task, error) {
	var tasks []*models.Task
	cursor, err := tasksCollection.Find(ctx, bson.D{})
	utils.HandleError(err)

	for cursor.Next(ctx) {
		var task models.Task
		err := cursor.Decode(&task)
		utils.HandleError(err)
		tasks = append(tasks, &task)
	}
	utils.HandleError(err)
	cursor.Close(ctx)
	return tasks, err
}

func CreateTask(task *models.Task) (*models.Task, error) {
	result, err := tasksCollection.InsertOne(ctx, task)
	utils.HandleError(err)
	task.ID = result.InsertedID.(primitive.ObjectID)
	return task, err
}

func UpdateTask(taskID string, task *models.Task) (*models.Task, error) {
	id, err := primitive.ObjectIDFromHex(taskID)
	utils.HandleError(err)
	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "title", Value: task.Title}, {Key: "description", Value: task.Description}}}}
	_, err = tasksCollection.UpdateOne(ctx, filter, update)
	utils.HandleError(err)
	task.ID = id
	return task, err
}

func CompleteTask(taskID string, task *models.Task) (*models.Task, error) {
	id, err := primitive.ObjectIDFromHex(taskID)
	utils.HandleError(err)
	filter1 := bson.D{{Key: "_id", Value: id}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "completed", Value: task.Completed}}}}
	_, err = tasksCollection.UpdateOne(ctx, filter1, update)
	utils.HandleError(err)
	var taskToReturn models.Task
	filter2 := bson.D{{Key: "_id", Value: id}}
	err = tasksCollection.FindOne(ctx, filter2).Decode(&taskToReturn)
	utils.HandleError(err)
	return &taskToReturn, err
}

func DeleteTask(taskID string) (*mongo.DeleteResult, error) {
	id, err := primitive.ObjectIDFromHex(taskID)
	utils.HandleError(err)
	filter := bson.D{{Key: "_id", Value: primitive.ObjectID(id)}}
	result, err := tasksCollection.DeleteOne(ctx, filter)
	utils.HandleError(err)
	return result, err
}
