package S3

import (
	"context"

	bucket "github.com/ogwurujohnson/bucket/lib"
	service "github.com/ogwurujohnson/bucket/lib/services"
)

type S3 struct {}

var _ service.Service = &S3{}

func (s *S3) Upload(ctx context.Context) (*bucket.Response, error) {
	return nil, nil
}

func (s *S3) Download(ctx context.Context) (*bucket.Response, error) {
	return nil, nil
}

func (s *S3) List(ctx context.Context) (*[]bucket.Response, error) {
	return nil, nil
}

func (s *S3) Delete(ctx context.Context) (bool, error) {
	return true, nil
}
