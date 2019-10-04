package db

import (
	"context"
	"time"

	"github.com/letrannhatviet/my_framework/db/types"
	"github.com/letrannhatviet/my_framework/sequence"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertStudent(req types.StudentAddReq) error {
	newID := sequence.GetNextID(Client.Database(dbName).Collection("counter"), "student_id_seq")
	student := types.Student{
		ID:        newID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		ClassName: req.ClassName,
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := Client.Database(dbName).Collection(dbCol).InsertOne(ctx, student)
	return err
}

func DeleteStudent(id int) (*mongo.DeleteResult, error) {
	res, err := Client.Database(dbName).Collection(dbCol).DeleteOne(context.TODO(), bson.D{{"id", id}})
	return res, err
}
