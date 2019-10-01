package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/letrannhatviet/my_framework/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	dbName = config.Config.MongoDB.Name
	dbCol  = "Student"
)

var Client *mongo.Client

func init() {
	connectionString := fmt.Sprintf(
		"mongodb://%s:%s@%s:%s",
		config.Config.MongoDB.User,
		config.Config.MongoDB.Password,
		config.Config.MongoDB.Host,
		config.Config.MongoDB.Port,
	)
	fmt.Println(connectionString)
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	defer cancel()
	Client = client
}
