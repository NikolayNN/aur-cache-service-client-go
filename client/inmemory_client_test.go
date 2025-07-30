package client

import (
	"context"
	"testing"

	cachesdk "multi-level-cache-go-sdk/cache"
)

func TestInMemoryClient(t *testing.T) {
	c := NewInMemoryCacheClient()
	ctx := context.Background()

	entry := cachesdk.CacheEntry[any]{CacheName: "users", Key: "1", Value: "Bob"}
	if err := c.PutAll(ctx, []cachesdk.CacheEntry[any]{entry}); err != nil {
		t.Fatalf("put: %v", err)
	}
	hits, err := c.GetAllAny(ctx, []cachesdk.CacheId{{CacheName: "users", Key: "1"}})
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	if len(hits) != 1 || !hits[0].Found || hits[0].Value != "Bob" {
		t.Fatalf("unexpected hits: %#v", hits)
	}
	if err := c.EvictAll(ctx, []cachesdk.CacheId{{CacheName: "users", Key: "1"}}); err != nil {
		t.Fatalf("evict: %v", err)
	}
	hits, _ = c.GetAllAny(ctx, []cachesdk.CacheId{{CacheName: "users", Key: "1"}})
	if hits[0].Found {
		t.Fatalf("expected miss after evict")
	}
}
