package utils

import (
	"context"
	"log"
	"os"
	"testing"
	"time"
)

func InTime(maxTime time.Duration, t *testing.T, todo func()) {
	ctx, cancel := context.WithTimeout(context.Background(), maxTime)
	defer cancel()

	done := make(chan struct{})
	go func() {
		todo()
		close(done)
	}()

	select {
	case <-ctx.Done():
		t.Fatalf("test timed out after %v", maxTime)
	case <-done:
	}
}

func ReadStringFromFile(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("failed to read file %s: %v", filename, err)
	}
	return string(data)
}
