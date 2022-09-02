package service

import (
	"context"

	bucket "github.com/ogwurujohnson/bucket/lib"
)

type Service interface {
	Upload(ctx context.Context) (*bucket.Response, error)
	Download(ctx context.Context) (*bucket.Response, error)
	List(ctx context.Context) (*[]bucket.Response, error)
	Delete(ctx context.Context) (bool, error)
}

type ServiceMappers struct {
	Bucket Service
}

func (s *ServiceMappers) Upload(ctx context.Context) (*bucket.Response, error) {
	return nil, nil
}

func (s *ServiceMappers) Download(ctx context.Context) (*bucket.Response, error) {
	return nil, nil
}

func (s *ServiceMappers) List(ctx context.Context) (*[]bucket.Response, error) {
	return nil, nil
}

func (s *ServiceMappers) Delete(ctx context.Context) (bool, error) {
	return false, nil
}