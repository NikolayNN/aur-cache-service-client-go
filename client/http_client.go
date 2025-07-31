package client

import (
	"context"

	cachesdk "github.com/nikolaynn/multi-level-cache-go-sdk/cache"
)

// HttpCacheClient wraps the SDK client.
type HttpCacheClient struct {
	sdk *cachesdk.Client
}

// NewHttpCacheClient creates a client with default options.
func NewHttpCacheClient(baseURL string) *HttpCacheClient {
	return &HttpCacheClient{sdk: cachesdk.New(baseURL)}
}

// NewHttpCacheClientWithThreshold sets the gzip threshold used by the SDK.
func NewHttpCacheClientWithThreshold(baseURL string, threshold int) *HttpCacheClient {
	return &HttpCacheClient{sdk: cachesdk.NewWithThreshold(baseURL, threshold)}
}

// GetAllAny fetches entries as interface values.
func (c *HttpCacheClient) GetAllAny(ctx context.Context, ids []cachesdk.CacheId) ([]cachesdk.CacheEntryHit[any], error) {
	return cachesdk.GetAll[any](c.sdk, ctx, ids)
}

// PutAll stores multiple entries.
func (c *HttpCacheClient) PutAll(ctx context.Context, entries []cachesdk.CacheEntry[any]) error {
	return c.sdk.PutAll(ctx, entries)
}

// EvictAll removes multiple entries.
func (c *HttpCacheClient) EvictAll(ctx context.Context, ids []cachesdk.CacheId) error {
	return c.sdk.EvictAll(ctx, ids)
}
