package S3

import (
	"context"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// rename to s3Response or s3
type Response struct {
	bucket string
	key    string
}

type S3 interface {
	Upload() (*Response, error)
	Download() (*Response, error)
	List() (*[]Response, error)
	Delete() (bool, error)
}

// convert to receiver method, so we can use the bucket and key in 
// in the library, repeat for others

// create a config file, where we have WithMethods that initializes bucket and key
// example: https://github.com/gocardless/gocardless-pro-go/blob/master/options.go#L53
// https://github.com/aws/aws-sdk-go
func Upload(ctx context.Context) *Response {
	sess := session.Must(session.NewSession())
	svc := s3.New(sess)

	_, err := svc.PutObjectWithContext(ctx, )
	
	return &Response{
		bucket: "bucket name",
		key:    "key",
	}
}

func Download() *Response {
	return &Response{
		bucket: "bucket name",
		key:    "key",
	}
}

func List() *[]Response {
	return &[]Response{
		{
			bucket: "bucket name",
			key:    "key",
		},
	}
}

func Delete() bool {
	return true
}
