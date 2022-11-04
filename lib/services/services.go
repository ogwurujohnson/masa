package services

import (
	"context"

	"github.com/ogwurujohnson/masa/lib/operation"
	"github.com/ogwurujohnson/masa/lib/services/gcs"
	"github.com/ogwurujohnson/masa/lib/services/s3"
)

type ServiceMappers struct {
	adapter operation.Operation
	bucket  string
	key     string
}

type InitType struct {
	AdapterType string
	Bucket      string
	Key         string
}


var DEFAULT_PAGE_SIZE int64 = 1000


func generateS3Config() *s3.S3 {
	config := &s3.Config{
		MaxRetries: 3,
		Region:     "us-west-2",
	}

	return s3.Build(config)
}

func generateGcsConfig() *gcs.Gcs {
	timeout := uint64(50)
	return gcs.Build(&timeout)
}

func Initialize(initData InitType) *ServiceMappers {
	switch initData.AdapterType {
	case "gcs":
		{
			return &ServiceMappers{
				adapter: generateGcsConfig(),
				bucket:  initData.Bucket,
				key:     initData.Key,
			}
		}
	case "s3":
		{
			return &ServiceMappers{
				adapter: generateS3Config(),
				bucket:  initData.Bucket,
				key:     initData.Key,
			}
		}
	default:
		{
			return &ServiceMappers{
				adapter: generateS3Config(),
				bucket:  initData.Bucket,
				key:     initData.Key,
			}
		}
	}
}

func (s *ServiceMappers) Upload(ctx context.Context, content any) (*operation.OperationResponse, error) {
	resp, err := s.adapter.Upload(ctx, s.bucket, s.key, content)
	return resp, err
}

func (s *ServiceMappers) Download(ctx context.Context) (*operation.OperationResponse, error) {
	resp, err := s.adapter.Download(ctx, s.bucket, s.key)
	return resp, err
}

func (s *ServiceMappers) List(ctx context.Context, pageSize *int64) (*operation.OperationResponse, error) {
	if pageSize == nil {
		pageSize = &DEFAULT_PAGE_SIZE
	}
	resp, err := s.adapter.List(ctx, s.bucket, s.key, *pageSize)
	return resp, err
}

func (s *ServiceMappers) Delete(ctx context.Context) (bool, error) {
	resp, err := s.adapter.Delete(ctx, s.bucket, s.key)
	return resp, err
}
