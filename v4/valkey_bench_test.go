package valkey

import (
	"context"
	"fmt"
	"math"
	"testing"

	lib_store "github.com/eko/gocache/lib/v4/store"
	vk "github.com/valkey-io/valkey-go"
	"github.com/valkey-io/valkey-go/valkeycompat"
)

func BenchmarkValkeySet(b *testing.B) {
	ctx := context.Background()

	vkClient, err := vk.NewClient(vk.ClientOption{
		InitAddress: []string{"localhost:6379"},
	})
	if err != nil {
		b.Fatal(err)
	}

	store := NewValkey(valkeycompat.NewAdapter(vkClient))

	for k := 0.; k <= 10; k++ {
		n := int(math.Pow(2, k))
		b.Run(fmt.Sprintf("%d", n), func(b *testing.B) {
			for i := 0; i < b.N*n; i++ {
				key := fmt.Sprintf("test-%d", n)
				value := []byte(fmt.Sprintf("value-%d", n))

				store.Set(ctx, key, value, lib_store.WithTags([]string{fmt.Sprintf("tag-%d", n)}))
			}
		})
	}
}

func BenchmarkValkeyGet(b *testing.B) {
	ctx := context.Background()

	vkClient, err := vk.NewClient(vk.ClientOption{
		InitAddress: []string{"localhost:6379"},
	})
	if err != nil {
		b.Fatal(err)
	}

	store := NewValkey(valkeycompat.NewAdapter(vkClient))

	key := "test"
	value := []byte("value")

	store.Set(ctx, key, value)

	for k := 0.; k <= 10; k++ {
		n := int(math.Pow(2, k))
		b.Run(fmt.Sprintf("%d", n), func(b *testing.B) {
			for i := 0; i < b.N*n; i++ {
				_, _ = store.Get(ctx, key)
			}
		})
	}
}
