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

func GetVan(ctx context.Context, targetId string) *models.ReturnVansStruct {
	objectID, err := primitive.ObjectIDFromHex(targetId)
	if err != nil {
		fmt.Println(err)
	}
	cursor := database.VanCollection.FindOne(ctx, bson.M{"_id": objectID})
	if cursor != nil {
		var result *models.ReturnVansStruct
		cursor.Decode(&result)
		return result
	}
	return nil
}

func GetAllVans(ctx context.Context) []*models.ReturnVansStruct {
	var results []*models.ReturnVansStruct
	cursor, err := database.VanCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(ctx) {
		var result models.ReturnVansStruct
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal("GetAllVans Decoder: ", err)
		}
		results = append(results, &result)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	return results
}

func CreateVan(ctx context.Context, targetVan *models.VansStruct) *mongo.InsertOneResult {
	result, err := database.VanCollection.InsertOne(ctx, targetVan)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func UpdateVan(ctx context.Context, vanId string, targetVan *models.VansStruct) int {
	objectID, err := primitive.ObjectIDFromHex(vanId)
	if err != nil {
		fmt.Println(err)
	}
	filter := bson.D{{Key: "_id", Value: objectID}}
	var updateData bson.D
	if targetVan.Name != "" {
		nameBson := bson.D{{Key: "name", Value: targetVan.Name}}
		updateData = append(updateData, nameBson...)
	}
	if targetVan.Desc != "" {
		descBson := bson.D{{Key: "desc", Value: targetVan.Desc}}
		updateData = append(updateData, descBson...)
	}
	if targetVan.Code != "" {
		codeBson := bson.D{{Key: "code", Value: targetVan.Code}}
		updateData = append(updateData, codeBson...)
	}
	update := bson.D{{Key: "$set", Value: updateData}}
	result, err := database.VanCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal(err)
	}
	return int(result.ModifiedCount)
}

func DeleteVan(ctx context.Context, targetId string) int {
	objectID, err := primitive.ObjectIDFromHex(targetId)
	if err != nil {
		fmt.Println(err)
	}
	result, err := database.VanCollection.DeleteOne(ctx, bson.D{{Key: "_id", Value: objectID}})
	if err != nil {
		log.Fatal(err)
	}
	return int(result.DeletedCount)
}
