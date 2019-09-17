package db

import (
	"context"

	"github.com/letrannhatviet/my_framework/db/types"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetStudent() (*types.Student, error) {
	var student types.Student
	err := collection.FindOne(context.TODO(), struct{}{}).Decode(&student)

	if err != nil {
		return nil, err
	}
	return &student, nil
}

func GetAllStudent() (*[]types.Student, error) {
	var students []types.Student
	findOptions := options.Find()
	findOptions.SetLimit(30)

	cur, err := collection.Find(context.TODO(), struct{}{})

	if err != nil {
		return nil, err
	}

	for cur.Next(context.TODO()) {
		var student types.Student
		err = cur.Decode(&student)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	return &students, nil
}
