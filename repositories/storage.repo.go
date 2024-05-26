package repositories

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"van_thailand_server/storage"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func UploadFileToS3(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	key := fileHeader.Filename
	log.Printf("uploading %q .....\n", key)
	_, err := storage.S3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(storage.BucketName),
		Key:    aws.String(key),
		Body:   file,
	})
	if err != nil {
		log.Printf("unable to upload %q, %v", key, err)
		return "", err
	}
	log.Printf("Successfully uploaded %q\n", key)
	url := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", storage.BucketName, key)
	return url, nil
}
