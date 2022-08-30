package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)

	database := client.Database("demo_with_go")
	podcastsCollection := database.Collection("test_go")

	//	findQuery := bson.D{{"$match", {"duration": 32}}}
	Cursor, err := podcastsCollection.Aggregate(ctx, mongo.Pipeline{{{"$match", bson.D{{"duration", bson.D{{"$gte", 20}}}}}}})
	if err != nil {
		panic(err)
	}
	var showsLoaded []bson.M
	if err = Cursor.All(ctx, &showsLoaded); err != nil {
		panic(err)
	}
	fmt.Println(showsLoaded)
}
