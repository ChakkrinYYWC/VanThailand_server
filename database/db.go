package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Database *mongo.Database

func ConnectDB(ctx context.Context) *mongo.Client {
	fmt.Println("Connecting to MongoDB...")
	var (
		mongoURI     = "mongodb+srv://forthAdmin:0807800687forth@vanthailand.9mwh1za.mongodb.net/"
		databaseName = "VanThailand"
	)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	Database = client.Database(databaseName)
	VanCollection = Database.Collection("Vans")
	VanScheduleCollection = Database.Collection("VanSchedule")
	UserCollection = Database.Collection("Users")
	fmt.Println("Connected to MongoDB!")
	return client
}
