package S3

import (
	"context"

	"github.com/ogwurujohnson/bucket/lib/services"
)

type S3 struct {
	// Config
}

var _ services.Service = &S3{}

func (s *S3) Upload(ctx context.Context) (*services.Response, error) {
	return nil, nil
}

func (s *S3) Download(ctx context.Context) (*services.Response, error) {
	return nil, nil
}

func (s *S3) List(ctx context.Context) (*[]services.Response, error) {
	return nil, nil
}

func (s *S3) Delete(ctx context.Context) (bool, error) {
	return true, nil
}
