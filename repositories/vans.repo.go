package repositories

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
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

func CreateVan(ctx context.Context, name string, code string, desc string, vanImage []string) *mongo.InsertOneResult {
	var updateData bson.D
	if name != "" {
		nameBson := bson.D{{Key: "name", Value: name}}
		updateData = append(updateData, nameBson...)
	}
	if code != "" {
		codeBson := bson.D{{Key: "code", Value: code}}
		updateData = append(updateData, codeBson...)
	}
	if desc != "" {
		descBson := bson.D{{Key: "desc", Value: desc}}
		updateData = append(updateData, descBson...)
	}
	if vanImage != nil {
		imageBson := bson.D{{Key: "imagePath", Value: vanImage}}
		updateData = append(updateData, imageBson...)
	}
	result, err := database.VanCollection.InsertOne(ctx, updateData)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func UpdateVan(ctx context.Context, vanId string, name string, code string, desc string, vanImages []string, imagePosition string) int {
	objectID, err := primitive.ObjectIDFromHex(vanId)
	if err != nil {
		fmt.Println(err)
	}
	filter := bson.D{{Key: "_id", Value: objectID}}
	var updateData bson.D
	if name != "" {
		nameBson := bson.D{{Key: "name", Value: name}}
		updateData = append(updateData, nameBson...)
	}
	if desc != "" {
		descBson := bson.D{{Key: "desc", Value: desc}}
		updateData = append(updateData, descBson...)
	}
	if code != "" {
		codeBson := bson.D{{Key: "code", Value: code}}
		updateData = append(updateData, codeBson...)
	}
	if vanImages != nil {
		var targetVan *models.ReturnVansStruct
		err = database.VanCollection.FindOne(ctx, filter).Decode(&targetVan)
		if err != nil {
			log.Fatal(err)
		}
		var newVanImages []string
		updatePosition := strings.Split(imagePosition, ",")
		que := 0
		for index := 0; index < 5; index++ {
			if que >= len(updatePosition) {
				break
			}
			positionInt, err := strconv.Atoi(updatePosition[que])
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("before_2")
			if index == positionInt {
				newVanImages = append(newVanImages, vanImages[que])
				que++
			} else {
				newVanImages = append(newVanImages, targetVan.ImagePath[index])
			}
		}
		imageBson := bson.D{{Key: "imagePath", Value: newVanImages}}
		updateData = append(updateData, imageBson...)
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
