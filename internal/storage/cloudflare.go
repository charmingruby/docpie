package storage

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/sirupsen/logrus"
)

func New(logger *logrus.Logger) *Cloudflare {
	creds := loadCredentials()

	r2Resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL: fmt.Sprintf("https://%s.r2.cloudflarestorage.com", creds.CloudflareAccountID),
		}, nil
	})

	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithEndpointResolverWithOptions(r2Resolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(creds.AccessKeyID, creds.AccessKeySecret, "")),
		config.WithRegion("auto"),
	)
	if err != nil {
		log.Fatal(err)
	}

	client := s3.NewFromConfig(cfg)

	return &Cloudflare{
		client:      client,
		credentials: creds,
		logger:      logger,
	}
}

type Cloudflare struct {
	client      *s3.Client
	credentials *cloudflareCredentials
	logger      *logrus.Logger
}

type cloudflareCredentials struct {
	BucketName          string
	CloudflareAccountID string
	AccessKeyID         string
	AccessKeySecret     string
}

func loadCredentials() *cloudflareCredentials {
	bucketName := os.Getenv("AWS_BUCKET_NAME")
	accountId := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	accessKeyId := os.Getenv("AWS_ACCESS_KEY_ID")
	accessKeySecret := os.Getenv("AWS_SECRET_ACCESS_KEY")

	return &cloudflareCredentials{
		BucketName:          bucketName,
		CloudflareAccountID: accountId,
		AccessKeyID:         accessKeyId,
		AccessKeySecret:     accessKeySecret,
	}
}
