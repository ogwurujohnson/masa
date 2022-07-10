package Gcs

import (
	"context"

	"cloud.google.com/go/storage"
)

type Response struct {
	bucket string
	key    string
}

type Gcs interface {
	Upload() (*Response, error)
	Download() (*Response, error)
	List() (*[]Response, error)
	Delete() (bool, error)
}

func Upload(ctx context.Context) (*Response, error) {
	_, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}

	return &Response{
		bucket: "bucket name",
		key:    "key",
	}, nil
}

func Download(ctx context.Context) (*Response, error) {
	_, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}

	return &Response{
		bucket: "bucket name",
		key:    "key",
	}, nil
}

func List(ctx context.Context) (*[]Response, error) {
	_, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}

	return &[]Response{
		{
			bucket: "bucket name",
			key:    "key",
		},
	}, nil
}

func Delete(ctx context.Context) (bool, error) {
	_, err := storage.NewClient(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}
