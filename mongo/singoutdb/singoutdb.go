package singoutdb

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func Getcol() *mongo.Collection {
	return collection
}

var ctx = context.TODO()

var server = "mongodb://34.101.230.112:27017"
var db_mongo = "auth_singout_demo_07_21"
var collec = "users_singout_demo_07_21"

func init() {
	clientOptions := options.Client().ApplyURI(server)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database(db_mongo).Collection(collec)
	fmt.Println("singoutdb ready to use ...")
}
