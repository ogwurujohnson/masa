package S3

import (
	"context"
	"io"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/ogwurujohnson/bucket/lib/services"
)

type S3 struct {
	timeout *uint64
	storage *s3.S3
}

type Config struct {
	MaxRetries     int
	Region         string
	timeoutSeconds *uint64
}

var _ services.Service = &S3{}

const (
	DefaultTimeoutSeconds uint64 = 30
)

func initialize(session *session.Session, timeoutSeconds *uint64) *S3 {
	client := s3.New(session)

	return &S3{
		storage: client,
	}
}

func Build(config *Config) *S3 {
	var initSession *session.Session
	if config == nil {
		initSession = session.Must(session.NewSession())
	} else {
		initSession = session.Must(session.NewSession(&aws.Config{
			Region:     aws.String(config.Region),
			MaxRetries: aws.Int(config.MaxRetries),
		}))
	}

	var timeout uint64 = DefaultTimeoutSeconds
	if config.timeoutSeconds != nil {
		return initialize(initSession, config.timeoutSeconds)
	}
	return initialize(initSession, &timeout)
}

func (s *S3) Upload(ctx context.Context, bucket string, key string, content interface{}) (*services.Response, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(*s.timeout))
	defer cancel()

	_, err := s.storage.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   content.(io.ReadSeeker),
	})

	if err != nil {
		return nil, err
	}

	return &services.Response{
		Bucket: bucket,
		Key:    key,
	}, nil
}

func (s *S3) Download(ctx context.Context, bucket string, key string) (*services.Response, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(*s.timeout))
	defer cancel()

	file, err := s.storage.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})

	if err != nil {
		return nil, err
	}

	return &services.Response{
		Bucket:  bucket,
		Key:     key,
		Content: file,
	}, nil
}

func (s *S3) List(ctx context.Context, bucket string, key string, pageSize int64) (*services.Response, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(*s.timeout))
	defer cancel()

	page, err := s.storage.ListObjectsV2(&s3.ListObjectsV2Input{
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
	return &services.Response{
		Bucket: bucket,
		Keys:   mapped,
	}, nil
}

func (s *S3) Delete(ctx context.Context, bucket string, key string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(*s.timeout))
	defer cancel()
	
	_, err := s.storage.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})

	if err != nil {
		return false, err
	}

	return true, nil
}
