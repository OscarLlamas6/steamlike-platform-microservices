package utils

import (
	"bytes"
	"developers-service/models"
	"encoding/base64"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"
)

func UploadImage(image string) (string, bool) {

	// INICIANDO A SUBIR IMAGEN A BUCKET S3
	id := uuid.New()
	myID := id.String()

	S3_ACCESS_KEY := os.Getenv("S3_ACCESS_KEY")
	S3_SECRET_KEY := os.Getenv("S3_SECRET_KEY")
	S3_AWS_REGION := os.Getenv("S3_AWS_REGION")
	S3_BUCKET_NAME := os.Getenv("S3_BUCKET_NAME")

	myAWSConfig := models.AWSConfig{
		AccessKeyID:     S3_ACCESS_KEY,
		AccessKeySecret: S3_SECRET_KEY,
		Region:          S3_AWS_REGION,
		BucketName:      S3_BUCKET_NAME,
		UploadTimeout:   100,
		BaseURL:         "",
	}

	mySession := CreateSession(myAWSConfig)

	imagenName := "sa_grupo4/" + myID + ".jpg"
	imagenURL := "https://" + myAWSConfig.BucketName + ".s3." + myAWSConfig.Region + ".amazonaws.com/" + imagenName

	myFinalBase64 := image[strings.IndexByte(image, ',')+1:]

	decode, errEnc := base64.StdEncoding.DecodeString(myFinalBase64)
	if errEnc != nil {
		fmt.Printf("Error: %v\n", errEnc)
		return "", false
	}

	uploader := s3manager.NewUploader(mySession)
	_, errS3 := uploader.Upload(&s3manager.UploadInput{
		Bucket:             &myAWSConfig.BucketName,
		Key:                &imagenName,
		Body:               bytes.NewReader(decode),
		ContentType:        aws.String("image/jpg"),
		ContentDisposition: aws.String("inline;"),
	})

	if errS3 != nil {
		fmt.Printf("Error: %v\n", errS3)
		return "", false
	}

	return imagenURL, true
}
