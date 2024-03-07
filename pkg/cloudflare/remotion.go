package cloudflare

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (c *Cloudflare) Remotion(key string) error {
	c.logger.Info("Removing a file from Cloudflare...")

	_, err := c.client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(c.credentials.BucketName),
		Key:    aws.String(key),
	})
	if err != nil {
		c.logger.Info(fmt.Sprintf("Error removing Cloudflare file: %v", err))

		return err
	}

	c.logger.Info("Removed successfully the file from Cloudflare.")

	return nil
}
