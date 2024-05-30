package services

import (
	"context"
	"log"
	"mime/multipart"
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

func CreateVan(ctx context.Context, name string, code string, desc string, files []*multipart.FileHeader) *mongo.InsertOneResult {
	var vanImages []string
	for _, file := range files {
		openedFile, err := file.Open()
		if err != nil {
			log.Println("openedFile error: ", err)
			return nil
		}
		url, err := repositories.UploadFileToS3(openedFile, file)
		if err != nil {
			log.Println("url error: ", err)
			return nil
		}
		vanImages = append(vanImages, url)
		openedFile.Close()
	}
	result := repositories.CreateVan(ctx, name, code, desc, vanImages)
	return result
}

func UpdateVan(ctx context.Context, targetId string, name string, code string, desc string, files []*multipart.FileHeader, imagePosition string) int {
	var vanImages []string
	for _, file := range files {
		openedFile, err := file.Open()
		if err != nil {
			log.Println("openedFile error: ", err)
			return 0
		}
		url, err := repositories.UploadFileToS3(openedFile, file)
		if err != nil {
			log.Println("url error: ", err)
			return 0
		}
		vanImages = append(vanImages, url)
		openedFile.Close()
	}
	result := repositories.UpdateVan(ctx, targetId, name, code, desc, vanImages, imagePosition)
	return result
}

func DeleteVan(ctx context.Context, targetId string) int {
	result := repositories.DeleteVan(ctx, targetId)
	return result
}
