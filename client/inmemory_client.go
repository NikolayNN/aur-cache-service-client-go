package client

import (
	"context"
	"sync"

	cachesdk "github.com/nikolaynn/multi-level-cache-go-sdk/cache"
)

// InMemoryCacheClient stores entries in a map for testing.
type InMemoryCacheClient struct {
	mu    sync.RWMutex
	store map[cachesdk.CacheId]any
}

// NewInMemoryCacheClient creates an empty in-memory client.
func NewInMemoryCacheClient() *InMemoryCacheClient {
	return &InMemoryCacheClient{store: make(map[cachesdk.CacheId]any)}
}

// GetAllAny returns cached values for the given ids.
func (c *InMemoryCacheClient) GetAllAny(ctx context.Context, ids []cachesdk.CacheId) ([]cachesdk.CacheEntryHit[any], error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	hits := make([]cachesdk.CacheEntryHit[any], 0, len(ids))
	for _, id := range ids {
		val, ok := c.store[id]
		hits = append(hits, cachesdk.CacheEntryHit[any]{
			CacheName: id.CacheName,
			Key:       id.Key,
			Value:     val,
			Found:     ok,
		})
	}
	return hits, nil
}

// PutAll stores entries in memory.
func (c *InMemoryCacheClient) PutAll(ctx context.Context, entries []cachesdk.CacheEntry[any]) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	for _, e := range entries {
		id := cachesdk.CacheId{CacheName: e.CacheName, Key: e.Key}
		c.store[id] = e.Value
	}
	return nil
}

// EvictAll removes entries from memory.
func (c *InMemoryCacheClient) EvictAll(ctx context.Context, ids []cachesdk.CacheId) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	for _, id := range ids {
		delete(c.store, id)
	}
	return nil
}
