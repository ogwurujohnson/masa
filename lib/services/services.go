package services

import (
	"context"
)

type Service interface {
	Upload(ctx context.Context, bucket string, key string, content interface{}) (*Response, error)
	Download(ctx context.Context, bucket string, key string) (*Response, error)
	List(ctx context.Context) (*[]Response, error)
	Delete(ctx context.Context, bucket string, key string) (bool, error)
}

type ServiceMappers struct {
	Bucket Service
}

func (s *ServiceMappers) Upload(ctx context.Context) (*Response, error) {
	return nil, nil
}

func (s *ServiceMappers) Download(ctx context.Context) (*Response, error) {
	return nil, nil
}

func (s *ServiceMappers) List(ctx context.Context) (*[]Response, error) {
	return nil, nil
}

func (s *ServiceMappers) Delete(ctx context.Context) (bool, error) {
	return false, nil
}
