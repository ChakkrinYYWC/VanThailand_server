package repositories

import (
	"context"
	"van_thailand_server/database"
	"van_thailand_server/models"

	"go.mongodb.org/mongo-driver/bson"
)

func FindUser(ctx context.Context, userReq *models.UserStruct) (*models.UserStruct, error) {
	var userData *models.UserStruct
	err := database.UserCollection.FindOne(ctx, bson.M{"username": userReq.Username}).Decode(&userData)
	if err != nil {
		// http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return nil, err
	}
	return userData, nil
}

func CreateUser(ctx context.Context, user *models.UserStruct) error {
	_, err := database.UserCollection.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
