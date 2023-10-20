package helpers

import (
	"context"
	"log"
	"mime/multipart"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	awsRegion  = "us-east-1" // Update with your AWS region
	bucketName = "hannonapp" // Update with your AWS S3 bucket name
	uploadPath = "images/"   // Update with the desired upload path in your S3 bucket
)

type ClientUploader struct {
	s3Client   *s3.S3
	bucketName string
	uploadPath string
}

var Uploader *ClientUploader

func init() {
	// Initialize AWS session and S3 client
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(awsRegion),
	})

	if err != nil {
		log.Fatalf("Failed to create AWS session: %v", err)
	}

	Uploader = &ClientUploader{
		s3Client:   s3.New(sess),
		bucketName: bucketName,
		uploadPath: uploadPath,
	}
}

func (c *ClientUploader) UploadFile(file multipart.File, object string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	// Upload the file to AWS S3
	_, err := c.s3Client.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket: aws.String(c.bucketName),
		Key:    aws.String(c.uploadPath + object),
		Body:   file,
	})

	if err != nil {
		return err
	}

	return nil
}
