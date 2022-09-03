package gcs

import (
	"context"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/ogwurujohnson/bucket/lib/services"
)

type Gcs struct {
	// storage *s3.S3
}

var _ services.Service = &Gcs{}

func initialize(session *session.Session) *Gcs {
	return nil
}

func Build() *Gcs {
	return initialize(nil)
}

func (g *Gcs) Upload(ctx context.Context) (*services.Response, error) {
	return nil, nil
}

func (g *Gcs) Download(ctx context.Context) (*services.Response, error) {
	return nil, nil
}

func (g *Gcs) List(ctx context.Context) (*[]services.Response, error) {
	return nil, nil
}

func (g *Gcs) Delete(ctx context.Context) (bool, error) {
	return true, nil
}
