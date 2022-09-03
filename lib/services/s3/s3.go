package S3

import (
	"context"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/ogwurujohnson/bucket/lib/services"
)

type S3 struct {
	storage *s3.S3
}

var _ services.Service = &S3{}

func initialize(session *session.Session) *S3 {
	svc := s3.New(session)
	return &S3{
		storage: svc,
	}
}

func Build() *S3 {
	session := session.Must(session.NewSession())
	return initialize(session)
}

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
