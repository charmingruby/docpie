package cloudflare

import (
	"context"
	"fmt"
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
		c.logger.Info(fmt.Sprintf("Error uploading Cloudflare file: %v", err))

		return err
	}

	c.logger.Info("File uploaded successfully to Cloudflare.")

	return nil
}
