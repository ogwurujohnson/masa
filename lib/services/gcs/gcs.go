package gcs

import (
	"context"

	bucket "github.com/ogwurujohnson/bucket/lib"
	service "github.com/ogwurujohnson/bucket/lib/services"
)

type Gcs struct {}

var _ service.Service = &Gcs{}

func (s *Gcs) Upload(ctx context.Context) (*bucket.Response, error) {
	return nil, nil
}

func (s *Gcs) Download(ctx context.Context) (*bucket.Response, error) {
	return nil, nil
}

func (s *Gcs) List(ctx context.Context) (*[]bucket.Response, error) {
	return nil, nil
}

func (s *Gcs) Delete(ctx context.Context) (bool, error) {
	return true, nil
}
