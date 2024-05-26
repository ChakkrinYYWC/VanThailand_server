package storage

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var (
	S3Client   *s3.Client
	BucketName = "vanthailandimage"
)

func Init(ctx context.Context) {
	log.Println("Init AWS S3 configs...")
	if os.Getenv("AWS_ACCESS_KEY_ID") == "" || os.Getenv("AWS_SECRET_ACCESS_KEY") == "" {
		log.Fatal("AWS_KEY is not set.")
	}
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("ap-southeast-1"),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			os.Getenv("AWS_ACCESS_KEY_ID"),
			os.Getenv("AWS_SECRET_ACCESS_KEY"),
			"",
		)))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	S3Client = s3.NewFromConfig(cfg)
}
