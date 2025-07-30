package persistent

import (
	"context"
	"fmt"

	"aur-cache-service-client-go/models"
	cachesdk "multi-level-cache-go-sdk/cache"
)

// PersistanceCacheLastMessageState provides typed helpers for LastMessageStateCacheEntry values.
type PersistanceCacheLastMessageState struct {
	cache *PersistentCache
}

// NewPersistanceCacheLastMessageState wraps a PersistentCache.
func NewPersistanceCacheLastMessageState(c *PersistentCache) *PersistanceCacheLastMessageState {
	return &PersistanceCacheLastMessageState{cache: c}
}

// Get retrieves a single entry.
func (l *PersistanceCacheLastMessageState) Get(ctx context.Context, key int64) (models.LastMessageStateCacheEntry, bool, error) {
	return Get[models.LastMessageStateCacheEntry](l.cache, ctx, fmt.Sprint(key))
}

// GetAll retrieves multiple entries.
func (l *PersistanceCacheLastMessageState) GetAll(ctx context.Context, keys []int64) ([]cachesdk.CacheEntryHit[models.LastMessageStateCacheEntry], error) {
	strKeys := make([]string, len(keys))
	for i, k := range keys {
		strKeys[i] = fmt.Sprint(k)
	}
	return GetAll[models.LastMessageStateCacheEntry](l.cache, ctx, strKeys)
}

// Put stores a single value.
func (l *PersistanceCacheLastMessageState) Put(ctx context.Context, key int64, val models.LastMessageStateCacheEntry) error {
	return l.cache.Put(ctx, fmt.Sprint(key), val)
}

// PutAll stores multiple values.
func (l *PersistanceCacheLastMessageState) PutAll(ctx context.Context, vals map[int64]models.LastMessageStateCacheEntry) error {
	m := make(map[string]any, len(vals))
	for k, v := range vals {
		m[fmt.Sprint(k)] = v
	}
	return l.cache.PutAll(ctx, m)
}

// Evict removes a key.
func (l *PersistanceCacheLastMessageState) Evict(ctx context.Context, key int64) error {
	return l.cache.Evict(ctx, fmt.Sprint(key))
}

// EvictAll removes multiple keys.
func (l *PersistanceCacheLastMessageState) EvictAll(ctx context.Context, keys []int64) error {
	str := make([]string, len(keys))
	for i, k := range keys {
		str[i] = fmt.Sprint(k)
	}
	return l.cache.EvictAll(ctx, str)
}
