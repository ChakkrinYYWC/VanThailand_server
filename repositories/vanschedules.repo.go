package repositories

import (
	"context"
	"fmt"
	"log"
	"van_thailand_server/database"
	"van_thailand_server/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetVanSchedule(ctx context.Context, scheduleId string) *models.ReturnScheduleStruct {
	objectID, err := primitive.ObjectIDFromHex(scheduleId)
	if err != nil {
		fmt.Println(err)
	}
	filter := bson.M{"_id": objectID}
	cursor, err := database.VanScheduleCollection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	if cursor.Next(ctx) {
		var result *models.ReturnScheduleStruct
		err = cursor.Decode(&result)
		if err != nil {
			log.Fatal("GetVanSchedule Decoder: ", err)
		}
		return result
	}
	return nil
}

func GetVanSchedules(ctx context.Context, targetVanId string) []*models.ReturnScheduleStruct {
	fmt.Println(targetVanId)
	var results []*models.ReturnScheduleStruct
	cursor, err := database.VanScheduleCollection.Find(ctx, bson.M{"van_id": targetVanId})
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(ctx) {
		var result models.ReturnScheduleStruct
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal("GetVanSchedules Decoder: ", err)
		}
		results = append(results, &result)
	}
	fmt.Println(results)
	return results
}

func CreateSchedule(ctx context.Context, vanId string, date string, destination string) *mongo.InsertOneResult {
	var updateData bson.D
	if vanId != "" {
		vanIdBson := bson.D{{Key: "van_id", Value: vanId}}
		updateData = append(updateData, vanIdBson...)
	}
	if date != "" {
		dateBson := bson.D{{Key: "date", Value: date}}
		updateData = append(updateData, dateBson...)
	}
	if destination != "" {
		destinationBson := bson.D{{Key: "destination", Value: destination}}
		updateData = append(updateData, destinationBson...)
	}
	update := bson.D{{Key: "$set", Value: updateData}}
	result, err := database.VanScheduleCollection.InsertOne(ctx, update)
	if err != nil {
		log.Fatal("CreateSchedule: ", err)
	}
	return result
}

func UpdateSchedule(ctx context.Context, scheduleId string, vanId string, date string, destination string) int {
	objectID, err := primitive.ObjectIDFromHex(scheduleId)
	if err != nil {
		fmt.Println(err)
	}
	filter := bson.D{{Key: "_id", Value: objectID}}
	var updateData bson.D
	if vanId != "" {
		vanIdBson := bson.D{{Key: "van_id", Value: vanId}}
		updateData = append(updateData, vanIdBson...)
	}
	if date != "" {
		dateBson := bson.D{{Key: "date", Value: date}}
		updateData = append(updateData, dateBson...)
	}
	if destination != "" {
		destinationBson := bson.D{{Key: "destination", Value: destination}}
		updateData = append(updateData, destinationBson...)
	}
	update := bson.D{{Key: "$set", Value: updateData}}
	result, err := database.VanScheduleCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal(err)
	}
	return int(result.ModifiedCount)
}

func DeleteSchedule(ctx context.Context, scheduleId string) int {
	objectID, err := primitive.ObjectIDFromHex(scheduleId)
	if err != nil {
		fmt.Println(err)
	}
	result, err := database.VanScheduleCollection.DeleteOne(ctx, bson.D{{Key: "_id", Value: objectID}})
	if err != nil {
		log.Fatal(err)
	}
	return int(result.DeletedCount)
}
