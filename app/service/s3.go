package service

import (
	"context"
	// "fmt"
	"io"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/pkg/errors"
)

type S3Uploader interface {
	UploadFile(ctx context.Context, file io.Reader, fileName string) (s3URL string, err error)
}

type DefaultS3Uploader struct {
	client *s3.Client
	bucket string
}

func NewS3Uploader(client *s3.Client, bucket string) S3Uploader {
	return &DefaultS3Uploader{
		client: client,
		bucket: bucket,
	}
}

// UploadFile uploads a file to S3 and returns its URL.
func (u *DefaultS3Uploader) UploadFile(ctx context.Context, file io.Reader, fileName string) (string, error) {
	input := &s3.PutObjectInput{
		Bucket: aws.String(u.bucket),
		Key:    aws.String(fileName),
		Body:   file,
	}

	_, err := u.client.PutObject(ctx, input)
	if err != nil {
		return "", errors.Wrap(err, "failed to upload file to S3")
	}

	return fileName, nil
}
