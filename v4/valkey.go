package valkey

import (
	"context"
	"fmt"
	"time"

	lib_store "github.com/eko/gocache/lib/v4/store"
	"github.com/valkey-io/valkey-go/valkeycompat"
)

// ValkeyClientInterface represents a valkey-io/valkey-go/valkeycompat client
type ValkeyClientInterface interface {
	Get(ctx context.Context, key string) *valkeycompat.StringCmd
	TTL(ctx context.Context, key string) *valkeycompat.DurationCmd
	Expire(ctx context.Context, key string, expiration time.Duration) *valkeycompat.BoolCmd
	Set(ctx context.Context, key string, values any, expiration time.Duration) *valkeycompat.StatusCmd
	Del(ctx context.Context, keys ...string) *valkeycompat.IntCmd
	FlushAll(ctx context.Context) *valkeycompat.StatusCmd
	SAdd(ctx context.Context, key string, members ...any) *valkeycompat.IntCmd
	SMembers(ctx context.Context, key string) *valkeycompat.StringSliceCmd
}

const (
	// ValkeyType represents the storage type as a string value
	ValkeyType = "valkey"
	// ValkeyTagPattern represents the tag pattern to be used as a key in specified storage
	ValkeyTagPattern = "gocache_tag_%s"
)

// ValkeyStore is a store for Redis
type ValkeyStore struct {
	client  ValkeyClientInterface
	options *lib_store.Options
}

// NewValkey creates a new store to Redis instance(s)
func NewValkey(client ValkeyClientInterface, options ...lib_store.Option) *ValkeyStore {
	return &ValkeyStore{
		client:  client,
		options: lib_store.ApplyOptions(options...),
	}
}

// Get returns data stored from a given key
func (s *ValkeyStore) Get(ctx context.Context, key any) (any, error) {
	object, err := s.client.Get(ctx, key.(string)).Result()
	if err == valkeycompat.Nil {
		return nil, lib_store.NotFoundWithCause(err)
	}
	return object, err
}

// GetWithTTL returns data stored from a given key and its corresponding TTL
func (s *ValkeyStore) GetWithTTL(ctx context.Context, key any) (any, time.Duration, error) {
	object, err := s.client.Get(ctx, key.(string)).Result()
	if err == valkeycompat.Nil {
		return nil, 0, lib_store.NotFoundWithCause(err)
	}
	if err != nil {
		return nil, 0, err
	}

	ttl, err := s.client.TTL(ctx, key.(string)).Result()
	if err != nil {
		return nil, 0, err
	}

	return object, ttl, err
}

// Set defines data in Redis for given key identifier
func (s *ValkeyStore) Set(ctx context.Context, key any, value any, options ...lib_store.Option) error {
	opts := lib_store.ApplyOptionsWithDefault(s.options, options...)

	err := s.client.Set(ctx, key.(string), value, opts.Expiration).Err()
	if err != nil {
		return err
	}

	if tags := opts.Tags; len(tags) > 0 {
		s.setTags(ctx, key, tags)
	}

	return nil
}

func (s *ValkeyStore) setTags(ctx context.Context, key any, tags []string) {
	for _, tag := range tags {
		tagKey := fmt.Sprintf(ValkeyTagPattern, tag)
		s.client.SAdd(ctx, tagKey, key.(string))
		s.client.Expire(ctx, tagKey, 720*time.Hour)
	}
}

// Delete removes data from Redis for given key identifier
func (s *ValkeyStore) Delete(ctx context.Context, key any) error {
	_, err := s.client.Del(ctx, key.(string)).Result()
	return err
}

// Invalidate invalidates some cache data in Redis for given options
func (s *ValkeyStore) Invalidate(ctx context.Context, options ...lib_store.InvalidateOption) error {
	opts := lib_store.ApplyInvalidateOptions(options...)

	if tags := opts.Tags; len(tags) > 0 {
		for _, tag := range tags {
			tagKey := fmt.Sprintf(ValkeyTagPattern, tag)
			cacheKeys, err := s.client.SMembers(ctx, tagKey).Result()
			if err != nil {
				continue
			}

			for _, cacheKey := range cacheKeys {
				s.Delete(ctx, cacheKey)
			}

			s.Delete(ctx, tagKey)
		}
	}

	return nil
}

// GetType returns the store type
func (s *ValkeyStore) GetType() string {
	return ValkeyType
}

// Clear resets all data in the store
func (s *ValkeyStore) Clear(ctx context.Context) error {
	if err := s.client.FlushAll(ctx).Err(); err != nil {
		return err
	}

	return nil
}
