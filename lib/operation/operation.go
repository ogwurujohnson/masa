package operation

import "context"

type OperationResponse struct {
	Bucket  string
	Key     string
	Keys    []string
	Content interface{}
}

type Operation interface {
	Upload(ctx context.Context, bucket string, key string, content interface{}) (*OperationResponse, error)
	Download(ctx context.Context, bucket string, key string) (*OperationResponse, error)
	List(ctx context.Context, bucket string, key string, pageSize int64) (*OperationResponse, error)
	Delete(ctx context.Context, bucket string, key string) (bool, error)
}
