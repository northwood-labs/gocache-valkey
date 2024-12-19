# Valkey adapter for [Gocache](https://github.com/eko/gocache/)

[Valkey](https://valkey.io) is an open-source fork of Redis, [supported by AWS](https://aws.amazon.com/elasticache/what-is-valkey/) and others. As a result, the Valkey adapter is nearly identical to the Redis adapter.

This project wraps the official [valkey-io/valkey-go](https://github.com/valkey-io/valkey-go) Go library to ensure that it leverages properly-supported Valkey code as opposed to Redis code that we _hope_ remains compatible.

## Installation

To begin working with the latest version of gocache, you can import the library in your project:

```bash
go get github.com/eko/gocache/lib/v4
```

and then, import the store(s) you want to use between all available ones:

```bash
go get github.com/northwood-labs/gocache-valkey/v4
```

If you run into any errors, please be sure to run `go mod tidy` to clean your go.mod file.

## Valkey

```go
// import (
//     "github.com/eko/gocache/lib/v4/cache"
//     vkCache "github.com/northwood-labs/gocache-valkey/v4"
//     vk "github.com/valkey-io/valkey-go"
//     "github.com/valkey-io/valkey-go/valkeycompat"
// )

valkeyClient, err := vk.NewClient(vk.ClientOption{
    InitAddress: []string{"localhost:6379"},
})
if err != nil {
    panic(err)
}

defer valkeyClient.Close()

valkeyStore := vkCache.NewValkey(valkeycompat.NewAdapter(valkeyClient))

cacheManager := cache.New[string](valkeyStore)
err := cacheManager.Set(ctx, "my-key", "my-value", store.WithExpiration(15*time.Second))
if err != nil {
    panic(err)
}

value, err := cacheManager.Get(ctx, "my-key")
switch err {
    case nil:
        fmt.Printf("Get the key '%s' from the valkey cache. Result: %s", "my-key", value)
    case valkeycompat.Nil:
        fmt.Printf("Failed to find the key '%s' from the valkey cache.", "my-key")
    default:
        fmt.Printf("Failed to get the value from the valkey cache with key '%s': %v", "my-key", err)
}
```

## Community

Please feel free to contribute on this library and do not hesitate to open an issue if you want to discuss about a feature.
