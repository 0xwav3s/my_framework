package db

import (
	"context"
	"log"

	"github.com/letrannhatviet/my_framework/db/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetOneStudent(std types.StudentReq) (*types.Student, error) {
	var student types.Student
	filter := bson.M{
		"first_name": std.FirstName,
		"last_name":  std.LastName,
		"email":      std.Email,
		"class_name": std.ClassName,
	}
	err := Client.Database(dbName).Collection(dbCol).FindOne(context.TODO(), filter).Decode(&student)

	if err != nil {
		return nil, err
	}
	return &student, nil
}

func GetById(id int) (*types.Student, error) {
	var student types.Student
	err := Client.Database(dbName).Collection(dbCol).FindOne(context.TODO(), bson.M{"id": id}).Decode(&student)

	if err != nil {
		return nil, err
	}
	return &student, nil
}

func GetAllStudents() (*[]types.Student, error) {
	var students []types.Student
	findOptions := options.Find()
	findOptions.SetLimit(30)

	cur, err := Client.Database(dbName).Collection(dbCol).Find(context.TODO(), struct{}{})

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

func GroupByLast(student types.Student) (interface{}, error) {
	opts := options.Aggregate()
	pipeline := bson.A{
		bson.M{"$match": bson.M{"last_name": student.LastName}},
		bson.M{"$group": bson.M{
			"_id":         student.LastName,
			"last_name":   bson.M{"$first": "$last_name"},
			"first_names": bson.M{"$push": "$first_name"},
			"ids":         bson.M{"$push": "$id"},
		},
		},
	}
	cur, err := Client.Database(dbName).Collection(dbCol).Aggregate(context.TODO(), pipeline, opts)
	if err != nil {
		return "", err
	}
	students := []bson.M{}
	err = cur.All(context.TODO(), &students)

	return students, nil
}

func SearchLikeStudent(std types.StudentSearchReq) (*[]types.Student, error) {
	var students []types.Student
	filter := bson.D{}
	if std.FirstName != "" {
		filter = append(filter, bson.E{"first_name", bson.M{"$regex": "^" + std.FirstName + ".*", "$options": "i"}})
	}
	if std.LastName != "" {
		filter = append(filter, bson.E{"last_name", bson.M{"$regex": "^" + std.LastName + ".*", "$options": "i"}})
	}
	if std.Email != "" {
		filter = append(filter, bson.E{"email", bson.M{"$regex": "^" + std.Email + ".*", "$options": "i"}})
	}
	if std.ClassName != "" {
		filter = append(filter, bson.E{"class_name", bson.M{"$regex": "^" + std.ClassName + ".*", "$options": "i"}})
	}
	cur, err := Client.Database(dbName).Collection(dbCol).Find(context.TODO(), filter)

	if err != nil {
		return nil, err
	}
	for cur.Next(context.TODO()) {
		var student types.Student
		err = cur.Decode(&student)
		students = append(students, student)
	}
	if err = cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())
	return &students, nil
}

func SearchStudent(std types.StudentReq) (*[]types.Student, error) {
	var students []types.Student
	findOptions := options.Find()
	findOptions.SetLimit(30)
	filter := bson.D{}
	if std.FirstName != "" {
		filter = append(filter, bson.E{"first_name", std.FirstName})
	}
	if std.LastName != "" {
		filter = append(filter, bson.E{"last_name", std.LastName})
	}
	if std.Email != "" {
		filter = append(filter, bson.E{"email", std.Email})
	}
	if std.ClassName != "" {
		filter = append(filter, bson.E{"class_name", std.ClassName})
	}
	cur, err := Client.Database(dbName).Collection(dbCol).Find(context.TODO(), filter)

	if err != nil {
		return nil, err
	}
	for cur.Next(context.TODO()) {
		var student types.Student
		err = cur.Decode(&student)
		students = append(students, student)
	}
	if err = cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())
	return &students, nil

}
