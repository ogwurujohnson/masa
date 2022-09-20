package gcs

import (
	"context"
	"io"
	"log"
	"time"

	"cloud.google.com/go/storage"
	"github.com/ogwurujohnson/bucket/lib/services"
	"google.golang.org/api/iterator"
)

type Gcs struct {
	timeout *uint64
	storage *storage.Client
}

var _ services.Service = &Gcs{}

const (
	DefaultTimeoutSeconds uint64 = 30
)

func initialize(timeoutSeconds *uint64) *Gcs {
	client, err := storage.NewClient(context.Background())
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
		return nil
	}

	return &Gcs{
		timeout: timeoutSeconds,
		storage: client,
	}
}

func Build(timeoutSeconds *uint64) *Gcs {
	var timeout uint64 = DefaultTimeoutSeconds
	if timeoutSeconds != nil {
		return initialize(timeoutSeconds)
	}

	return initialize(&timeout)
}

func (g *Gcs) Upload(ctx context.Context, bucket string, key string, content interface{}) (*services.Response, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(*g.timeout))
	defer cancel()

	objectWriter := g.getBucket(bucket).Object(key).NewWriter(ctx)
	if _, err := io.Copy(objectWriter, content.(io.ReadSeeker)); err != nil {
		return nil, err
	}

	if err := objectWriter.Close(); err != nil {
		return nil, err
	}

	return &services.Response{
		Bucket: bucket,
		Key:    key,
	}, nil
}

func (g *Gcs) Download(ctx context.Context, bucket string, key string) (*services.Response, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(*g.timeout))
	defer cancel()

	objectReader, err := g.getBucket(bucket).Object(key).NewReader(ctx)
	if err != nil {
		return nil, err
	}

	defer objectReader.Close()
	content, err := io.ReadAll(objectReader)
	if err != nil {
		return nil, err
	}

	return &services.Response{
		Bucket:  bucket,
		Key:     key,
		Content: content,
	}, nil
}

func (g *Gcs) List(ctx context.Context, bucket string, key string, pageSize int64) (*services.Response, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(*g.timeout))
	defer cancel()

	query := &storage.Query{Prefix: key}
	var keys []string

	it := g.getBucket(bucket).Objects(ctx, query)
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		keys = append(keys, attrs.Name)
	}

	return &services.Response{
		Bucket: bucket,
		Keys:   keys,
	}, nil
}

func (g *Gcs) Delete(ctx context.Context, bucket string, key string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(*g.timeout))
	defer cancel()

	if err := g.getBucket(bucket).Object(key).Delete(ctx); err != nil {
		return false, err
	}
	
	return true, nil
}

func (g *Gcs) getBucket(name string) *storage.BucketHandle {
	return g.storage.Bucket(name)
}
