package service

import (
	"context"
	"log"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type UploadService interface {
	PresignURL(ctx context.Context, key string) (string, error)
}

type uploadService struct {
	awsSess *session.Session
}

func NewUploadService() UploadService {
	key, exists := os.LookupEnv("DO_KEY")
	if !exists {
		log.Fatal("DigitalOcean spaces key not defined")
	}
	secret, exists := os.LookupEnv("DO_SECRET")
	if !exists {
		log.Fatal("DigitalOcean spaces secret not defined")
	}
	endpoint, exists := os.LookupEnv("DO_ENDPOINT")
	if !exists {
		log.Fatal("DigitalOcean spaces endpoint not defined")
	}
	region := strings.Split(endpoint, ".")[0]

	s3Config := &aws.Config{
		Credentials: credentials.NewStaticCredentials(key, secret, ""),
		Endpoint:    aws.String(endpoint),
		Region:      aws.String(region),
	}

	newSession, err := session.NewSession(s3Config)
	if err != nil {
		log.Fatal("Failed to create s3 session in Upload Service")
	}

	return &uploadService{awsSess: newSession}
}

func (as *uploadService) PresignURL(ctx context.Context, key string) (string, error) {
	s3Client := s3.New(as.awsSess)

	req, _ := s3Client.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String("stories"),
		Key:    aws.String(key),
	})

	urlStr, header, err := req.PresignRequest(10 * time.Minute)
	if err != nil {
		log.Print(err)
		return "", err
	}

	log.Println(header)

	return urlStr, nil
}
