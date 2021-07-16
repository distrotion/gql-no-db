package maindb

import (
	"context"
	_ "fmt"
	_ "log"

	_ "go.mongodb.org/mongo-driver/mongo"
	_ "go.mongodb.org/mongo-driver/mongo/options"
)

// var collection *mongo.Collection

// func Getcol() *mongo.Collection {
// 	return collection
// }

var ctx = context.TODO()

var server = "mongodb://localhost:9000/"
var db_mongo = "auth_clinic_demo_07_21"
var collec = "users_clinic_demo_07_21"

func init() {
	// clientOptions := options.Client().ApplyURI(server)
	// client, err := mongo.Connect(ctx, clientOptions)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = client.Ping(ctx, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// collection = client.Database(db_mongo).Collection(collec)
	// fmt.Println("maindb ready to use ...")
}
