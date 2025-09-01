package persistent

import (
	"context"
	"encoding/json"
	"fmt"

	cachesdk "github.com/nikolaynn/multi-level-cache-go-sdk/cache"
)

// PersistentCache wraps a named cache managed by the service.
type PersistentCache struct {
	name   string
	client interface {
		GetAllAny(ctx context.Context, ids []cachesdk.CacheId) ([]cachesdk.CacheEntryHit[any], error)
		PutAll(ctx context.Context, entries []cachesdk.CacheEntry[any]) error
		EvictAll(ctx context.Context, ids []cachesdk.CacheId) error
	}
}

// NewPersistentCache creates a cache with the given name.
func NewPersistentCache(name string, c interface {
	GetAllAny(context.Context, []cachesdk.CacheId) ([]cachesdk.CacheEntryHit[any], error)
	PutAll(context.Context, []cachesdk.CacheEntry[any]) error
	EvictAll(context.Context, []cachesdk.CacheId) error
}) *PersistentCache {
	return &PersistentCache{name: name, client: c}
}

// Name returns the cache name.
func (c *PersistentCache) Name() string { return c.name }

// Put stores a single value.
func (c *PersistentCache) Put(ctx context.Context, key string, value any) error {
	entry := cachesdk.CacheEntry[any]{CacheName: c.name, Key: key, Value: value}
	return c.client.PutAll(ctx, []cachesdk.CacheEntry[any]{entry})
}

// PutAll stores multiple values.
func (c *PersistentCache) PutAll(ctx context.Context, values map[string]any) error {
	entries := make([]cachesdk.CacheEntry[any], 0, len(values))
	for k, v := range values {
		entries = append(entries, cachesdk.CacheEntry[any]{CacheName: c.name, Key: k, Value: v})
	}
	fmt.Printf("ENTRIES %+v\n", entries)
	return c.client.PutAll(ctx, entries)
}

// Get retrieves a value into the provided type.
func Get[T any](c *PersistentCache, ctx context.Context, key string) (T, bool, error) {
	hits, err := GetAll[T](c, ctx, []string{key})
	if err != nil || len(hits) == 0 {
		var zero T
		return zero, false, err
	}
	hit := hits[0]
	return hit.Value, hit.Found, nil
}

// GetAll fetches multiple values.
func GetAll[T any](c *PersistentCache, ctx context.Context, keys []string) ([]cachesdk.CacheEntryHit[T], error) {
	ids := make([]cachesdk.CacheId, 0, len(keys))
	for _, k := range keys {
		ids = append(ids, cachesdk.CacheId{CacheName: c.name, Key: k})
	}
	rawHits, err := c.client.GetAllAny(ctx, ids)
	if err != nil {
		return nil, err
	}
	hits := make([]cachesdk.CacheEntryHit[T], 0, len(rawHits))
	for _, h := range rawHits {
		v, ok := h.Value.(T)
		if !ok {
			var zero T
			if h.Value != nil {
				b, _ := json.Marshal(h.Value)
				json.Unmarshal(b, &zero)
				v = zero
			}
		}
		hits = append(hits, cachesdk.CacheEntryHit[T]{
			CacheName: h.CacheName,
			Key:       h.Key,
			Value:     v,
			Found:     h.Found,
		})
	}
	return hits, nil
}

// Evict removes a single key.
func (c *PersistentCache) Evict(ctx context.Context, key string) error {
	id := cachesdk.CacheId{CacheName: c.name, Key: key}
	return c.client.EvictAll(ctx, []cachesdk.CacheId{id})
}

// EvictAll removes multiple keys.
func (c *PersistentCache) EvictAll(ctx context.Context, keys []string) error {
	ids := make([]cachesdk.CacheId, 0, len(keys))
	for _, k := range keys {
		ids = append(ids, cachesdk.CacheId{CacheName: c.name, Key: k})
	}
	return c.client.EvictAll(ctx, ids)
}
