package service

import (
	"bytes"
	"context"
	"encoding/base64"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type UploadService interface {
	Upload(ctx context.Context, filename, base64File string) (string, error)
}

type uploadService struct {
	aws *session.Session
}

func NewUploadService() UploadService {
	endpoint := "fra1.digitaloceanspaces.com"
	region := "fra1"
	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint: &endpoint,
		Region:   &region,
	}))

	return &uploadService{aws: sess}
}

func (as *uploadService) Upload(ctx context.Context, filename, base64File string) (string, error) {
	decode, err := base64.StdEncoding.DecodeString(base64File)
	if err != nil {
		return "", err
	}

	uploader := s3manager.NewUploader(as.aws)

	res, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("avatars"),
		Key:    aws.String(filename),
		ACL:    aws.String("public-read"),
		Body:   bytes.NewReader(decode),
	})
	if err != nil {
		return "", err
	}

	return res.Location, nil
}
