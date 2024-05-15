package services

import (
	"context"
	"van_thailand_server/models"
	"van_thailand_server/repositories"

	"go.mongodb.org/mongo-driver/mongo"
)

func GetVan(ctx context.Context, targetId string) *models.ReturnVansStruct {
	result := repositories.GetVan(ctx, targetId)
	return result
}

func GetVans(ctx context.Context) []*models.ReturnVansStruct {
	targetVan := repositories.GetAllVans(ctx)
	return targetVan
}

func CreateVan(ctx context.Context, targetVan *models.VansStruct) *mongo.InsertOneResult {
	result := repositories.CreateVan(ctx, targetVan)
	return result
}

func UpdateVan(ctx context.Context, targetId string, targetVan *models.VansStruct) int {
	result := repositories.UpdateVan(ctx, targetId, targetVan)
	return result
}

func DeleteVan(ctx context.Context, targetId string) int {
	result := repositories.DeleteVan(ctx, targetId)
	return result
}
