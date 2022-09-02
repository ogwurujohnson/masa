package gcs

import (
	"context"

	"github.com/ogwurujohnson/bucket/lib/services"
)

type Gcs struct {
	// Config
}

var _ services.Service = &Gcs{}

func (s *Gcs) Upload(ctx context.Context) (*services.Response, error) {
	return nil, nil
}

func (s *Gcs) Download(ctx context.Context) (*services.Response, error) {
	return nil, nil
}

func (s *Gcs) List(ctx context.Context) (*[]services.Response, error) {
	return nil, nil
}

func (s *Gcs) Delete(ctx context.Context) (bool, error) {
	return true, nil
}
