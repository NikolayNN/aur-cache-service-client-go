# AUR Cache Service Client (Go)

This project provides Go utilities for working with the [multi-level cache service](https://github.com/NikolayNN/multi-level-cache-go-sdk). It mirrors the features of the Java helper library and exposes a simple API for working with named persistent caches.

## Usage

Create an HTTP client pointing at your cache service and expose persistent caches using the manager:

```go
client := client.NewHttpCacheClient("http://localhost:8080")
manager := persistent.NewPersistentCacheManager(client)

users := manager.GetCache("users")
users.Put(context.Background(), "1", User{Name: "Alice"})
```

In-memory mode is also available for tests or when the service is not required:

```go
client := client.NewInMemoryCacheClient()
manager := persistent.NewPersistentCacheManager(client)
```

`PersistentCache` also exposes bulk operations:

```go
cache := manager.GetCache("users")
entries := map[string]User{"1": {Name: "Ann"}, "2": {Name: "Bob"}}
cache.PutAll(context.Background(), entries)
hits, _ := cache.GetAll[User](context.Background(), []string{"1", "2"})
cache.EvictAll(context.Background(), []string{"1", "2"})
```

## Development

Run the tests with:

```bash
go test ./...
```
