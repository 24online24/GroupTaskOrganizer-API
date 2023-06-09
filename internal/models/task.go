package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Task este structura pentru un task
type Task struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	RoomID      primitive.ObjectID `bson:"room_id" json:"room_id"` // ID-ul camerei din care face parte task-ul
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	Completed   bool               `bson:"completed" json:"completed"` 
}
