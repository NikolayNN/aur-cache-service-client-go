package persistent

import (
	"context"
	"sync"

	cachesdk "multi-level-cache-go-sdk/cache"
)

// CacheClient defines the methods used by the manager.
type CacheClient interface {
	GetAllAny(ctx context.Context, ids []cachesdk.CacheId) ([]cachesdk.CacheEntryHit[any], error)
	PutAll(ctx context.Context, entries []cachesdk.CacheEntry[any]) error
	EvictAll(ctx context.Context, ids []cachesdk.CacheId) error
}

// PersistentCacheManager creates and holds named caches.
type PersistentCacheManager struct {
	client CacheClient
	mu     sync.Mutex
	caches map[string]*PersistentCache
}

// NewPersistentCacheManager creates a manager for the given client.
func NewPersistentCacheManager(c CacheClient) *PersistentCacheManager {
	return &PersistentCacheManager{client: c, caches: make(map[string]*PersistentCache)}
}

// GetCache returns (and creates if needed) a named cache.
func (m *PersistentCacheManager) GetCache(name string) *PersistentCache {
	m.mu.Lock()
	defer m.mu.Unlock()
	if c, ok := m.caches[name]; ok {
		return c
	}
	c := NewPersistentCache(name, m.client)
	m.caches[name] = c
	return c
}

// CacheNames lists all created cache names.
func (m *PersistentCacheManager) CacheNames() []string {
	m.mu.Lock()
	defer m.mu.Unlock()
	names := make([]string, 0, len(m.caches))
	for n := range m.caches {
		names = append(names, n)
	}
	return names
}
