package aws

import (
	"blog_server/global"
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
)

// UploadFile reads from a file and puts the data into an object in a bucket.

func UploadFile(objectKey string, file io.Reader) error {
	// set credentials
	creds := credentials.NewStaticCredentialsProvider(global.Config.AWS.AccessKey, global.Config.AWS.SecretKey, "")

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(global.Config.AWS.Region),
		config.WithCredentialsProvider(creds))
	if err != nil {
		global.Logger.Error(err)
		return err
	}

	client := s3.NewFromConfig(cfg)

	// upload
	input := &s3.PutObjectInput{
		Bucket: aws.String(global.Config.AWS.Bucket),
		Key:    aws.String(objectKey),
		Body:   file,
	}

	_, err = client.PutObject(context.TODO(), input)
	if err != nil {
		global.Logger.Error(err)
		return err
	}

	return nil
}
