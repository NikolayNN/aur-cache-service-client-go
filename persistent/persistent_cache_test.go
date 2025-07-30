package persistent

import (
	"context"
	"testing"

	"aur-cache-service-client-go/client"
)

type testUser struct{ Name string }

func TestCachePutGet(t *testing.T) {
	c := client.NewInMemoryCacheClient()
	mgr := NewPersistentCacheManager(c)
	cache := mgr.GetCache("users")

	if err := cache.Put(context.Background(), "1", testUser{Name: "Alice"}); err != nil {
		t.Fatalf("put: %v", err)
	}

	u, found, err := Get[testUser](cache, context.Background(), "1")
	if err != nil || !found || u.Name != "Alice" {
		t.Fatalf("unexpected get: %v %v %#v", err, found, u)
	}
}
