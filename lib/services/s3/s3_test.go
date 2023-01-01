package s3_test

import (
	"context"
	"testing"

	"github.com/ogwurujohnson/masa/lib/services/s3"
)

func TestS3Upload(t *testing.T) {
	s := s3.Build(nil)
	ctx := context.TODO()

	// Test data
	bucket := "test-s3-bucket"
	key := "test.txt"
	body := []byte("This is a test.")

	_, err := s.Upload(ctx, bucket, key, body)
	if err != nil {
		t.Errorf("Failed to upload object %v", err)
	}
}
