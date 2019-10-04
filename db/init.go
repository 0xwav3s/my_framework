package db

import (
	"context"
	"fmt"
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
	connectionString := ""
	if config.Config.MongoDB.Protocol == "mongodb" {
		connectionString = fmt.Sprintf(
			"%s://%s:%s@%s:%s",
			config.Config.MongoDB.Protocol,
			config.Config.MongoDB.User,
			config.Config.MongoDB.Password,
			config.Config.MongoDB.Host,
			config.Config.MongoDB.Port,
		)
	} else {
		connectionString = fmt.Sprintf(
			"%s://%s:%s@%s",
			config.Config.MongoDB.Protocol,
			config.Config.MongoDB.User,
			config.Config.MongoDB.Password,
			config.Config.MongoDB.Host,
		)
	}
	connectionString = "mongodb+srv://mongoadmin:secret1234@cluster0-xxyrd.gcp.mongodb.net"

	fmt.Println(connectionString)
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}
	Client = client
}
