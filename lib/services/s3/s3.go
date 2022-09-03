package S3

import (
	"context"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/ogwurujohnson/bucket/lib/services"
)

type S3 struct {
	storage *s3.S3
}

type Config struct {
	MaxRetries int
	Region     string
}

var _ services.Service = &S3{}

func initialize(session *session.Session) *S3 {
	svc := s3.New(session)

	return &S3{
		storage: svc,
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

	return initialize(initSession)
}

func (s *S3) Upload(ctx context.Context, bucket string, key string, content interface{}) (*services.Response, error) {
	_, err := s.storage.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   content.(io.ReadSeeker),
	})

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok && aerr.Code() == request.CanceledErrorCode {
			fmt.Printf("upload canceled due to timeout,  %v\n", err)
			return nil, err
		}
		fmt.Printf("failed to upload object, %v\n", err)
		return nil, err
	}

	return &services.Response{
		Bucket: bucket,
		Key:    key,
	}, nil
}

func (s *S3) Download(ctx context.Context, bucket string, key string) (*services.Response, error) {
	file, err := s.storage.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key: aws.String(key),
	})

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok && aerr.Code() == request.CanceledErrorCode {
			fmt.Printf("upload canceled due to timeout,  %v\n", err)
			return nil, err
		}
		fmt.Printf("failed to upload object, %v\n", err)
		return nil, err
	}

	return &services.Response{
		Bucket: bucket,
		Key: key,
		Content: file,
	}, nil
}

func (s *S3) List(ctx context.Context) (*[]services.Response, error) {
	return nil, nil
}

func (s *S3) Delete(ctx context.Context) (bool, error) {
	return true, nil
}
