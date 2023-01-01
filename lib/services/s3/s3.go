package s3

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/ogwurujohnson/masa/lib/operation"
)

type S3 struct {
	timeout *uint64
	Storage *s3.S3
}

type Config struct {
	MaxRetries     int
	Region         string
	timeoutSeconds *uint64
}

var _ operation.Operation = &S3{}

const (
	DefaultTimeoutSeconds uint64 = 30
)

func initialize(session *session.Session, timeoutSeconds *uint64) *S3 {
	client := s3.New(session)

	return &S3{
		timeout: timeoutSeconds,
		Storage: client,
	}
}

func Build(config *Config) *S3 {
	var initSession *session.Session
	var timeout uint64 = DefaultTimeoutSeconds

	if config == nil {
		initSession = session.Must(session.NewSession())
		return initialize(initSession, &timeout)
	}

	initSession = session.Must(session.NewSession(&aws.Config{
		Region:     aws.String(config.Region),
		MaxRetries: aws.Int(config.MaxRetries),
	}))

	if config.timeoutSeconds != nil {
		timeout = *config.timeoutSeconds
	}

	return initialize(initSession, &timeout)
}

func (s *S3) Upload(ctx context.Context, bucket string, key string, content interface{}) (*operation.OperationResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(*s.timeout))
	defer cancel()

	_, err := s.Storage.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   content.(io.ReadSeeker),
	})

	if err != nil {
		return nil, err
	}

	return &operation.OperationResponse{
		Bucket: bucket,
		Key:    key,
	}, nil
}

func (s *S3) Download(ctx context.Context, bucket string, key string) (*operation.OperationResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(*s.timeout))
	defer cancel()

	file, err := s.Storage.GetObjectWithContext(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}
	defer file.Body.Close()

	content, err := io.ReadAll(file.Body)
	if err != nil {
		log.Printf("Failed to read downloaded object: %v", err)
		return nil, err
	}

	return &operation.OperationResponse{
		Bucket:  bucket,
		Key:     key,
		Content: content,
	}, nil
}

func (s *S3) List(ctx context.Context, bucket string, key string, pageSize int64) (*operation.OperationResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(*s.timeout))
	defer cancel()

	page, err := s.Storage.ListObjectsV2WithContext(ctx, &s3.ListObjectsV2Input{
		Prefix:  aws.String(key),
		Bucket:  aws.String(bucket),
		MaxKeys: aws.Int64(pageSize),
	})

	if err != nil {
		return nil, err
	}

	mapped := make([]string, *page.KeyCount)
	for i, objects := range page.Contents {
		mapped[i] = *objects.Key
	}
	return &operation.OperationResponse{
		Bucket: bucket,
		Keys:   mapped,
	}, nil
}

func (s *S3) Delete(ctx context.Context, bucket string, key string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(*s.timeout))
	defer cancel()

	_, err := s.Storage.DeleteObjectWithContext(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})

	if err != nil {
		return false, err
	}

	return true, nil
}
