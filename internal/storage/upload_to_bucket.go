package storage

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (c *Cloudflare) UploadToBucket(file io.Reader, key string) error {
	c.logger.Info("Uploading a file to Cloudflare...")

	_, err := c.client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(c.credentials.BucketName),
		Key:    aws.String(key),
		Body:   file,
	})
	if err != nil {
		return err
	}

	c.logger.Info("File uploaded successfully to Cloudflare.")

	return nil
}
