package database

import (
	"context"
	"fmt"
	configuration "golang-structure/src/configs"
	"log"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MgDB MongoInstance

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

func ConnectMongo() {
	client, _ := mongo.NewClient(options.Client().ApplyURI(configuration.Config.GetString("mongo.url")))

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := client.Connect(ctx)
	db := client.Database(configuration.Config.GetString("mongo.db_name"))

	if err != nil {
		fmt.Println(strings.Repeat("!", 40))
		fmt.Println("☹️  Could Not Establish Mongo DB Connection")
		fmt.Println(strings.Repeat("!", 40))

		log.Fatal(err)
	}

	fmt.Println(strings.Repeat("-", 40))
	fmt.Println("😀 Connected To Mongo DB")
	fmt.Println(strings.Repeat("-", 40))

	MgDB = MongoInstance{
		Client: client,
		Db:     db,
	}
}
