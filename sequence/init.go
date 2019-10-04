package sequence

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Sequence_Student struct {
	Key    string `json:"student_id_seq"`
	Number int    `json:"number"`
}

func GetNextID(col *mongo.Collection, seq string) int {
	var sequence Sequence_Student
	filter := bson.M{"key": seq}
	update := bson.M{"$inc": bson.M{"number": 1}}
	err := col.FindOneAndUpdate(context.TODO(), filter, update).Decode(&sequence)
	if err != nil {
		col.InsertOne(context.TODO(), Sequence_Student{seq, 0})
		fmt.Println(err)
	}
	return sequence.Number + 1
}
