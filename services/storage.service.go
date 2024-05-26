package services

import (
	"mime/multipart"
	"van_thailand_server/repositories"
)

func UploadFileToS3(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	url, err := repositories.UploadFileToS3(file, fileHeader)
	return url, err
}
