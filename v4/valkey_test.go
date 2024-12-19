package valkey

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/valkey-io/valkey-go/valkeycompat"
	"go.uber.org/mock/gomock"

	lib_store "github.com/eko/gocache/lib/v4/store"
)

func TestNewValkey(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)

	client := NewMockValkeyClientInterface(ctrl)

	// When
	store := NewValkey(client, lib_store.WithExpiration(6*time.Second))

	// Then
	assert.IsType(t, new(ValkeyStore), store)
	assert.Equal(t, client, store.client)
	assert.Equal(t, &lib_store.Options{Expiration: 6 * time.Second}, store.options)
}

func TestValkeyGet(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	client := NewMockValkeyClientInterface(ctrl)
	store := NewValkey(client)

	tests := []struct {
		name      string
		key       string
		returnVal string
		returnErr error
		expectErr bool
		expectVal interface{}
	}{
		{
			name:      "Returns Value",
			key:       "my-key",
			returnVal: "value",
			returnErr: nil,
			expectErr: false,
			expectVal: "value",
		},
		{
			name:      "Key Not Found",
			key:       "non-existent-key",
			returnVal: "",
			returnErr: valkeycompat.Nil,
			expectErr: true,
			expectVal: nil,
		},
		{
			name:      "Return Error",
			key:       "my-key",
			returnVal: "",
			returnErr: fmt.Errorf("some error"),
			expectErr: true,
			expectVal: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Given: mock the Valkey client's Get method
			client.EXPECT().Get(ctx, tt.key).Return(NewStringResult(tt.returnVal, tt.returnErr))

			// When
			value, err := store.Get(ctx, tt.key)

			// Then
			if tt.expectErr {
				assert.Error(t, err)
				assert.Equal(t, tt.expectVal, value)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.expectVal, value)
			}
		})
	}
}

func TestValkeySet(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)

	ctx := context.Background()

	cacheKey := "my-key"
	cacheValue := "my-cache-value"

	client := NewMockValkeyClientInterface(ctrl)
	client.EXPECT().Set(ctx, "my-key", cacheValue, 5*time.Second).Return(&valkeycompat.StatusCmd{})

	store := NewValkey(client, lib_store.WithExpiration(6*time.Second))

	// When
	err := store.Set(ctx, cacheKey, cacheValue, lib_store.WithExpiration(5*time.Second))

	// Then
	assert.Nil(t, err)
}

func TestValkeySetWhenNoOptionsGiven(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)

	ctx := context.Background()

	cacheKey := "my-key"
	cacheValue := "my-cache-value"

	client := NewMockValkeyClientInterface(ctrl)
	client.EXPECT().Set(ctx, "my-key", cacheValue, 6*time.Second).Return(&valkeycompat.StatusCmd{})

	store := NewValkey(client, lib_store.WithExpiration(6*time.Second))

	// When
	err := store.Set(ctx, cacheKey, cacheValue)

	// Then
	assert.Nil(t, err)
}

func TestValkeySetWithTags(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)

	ctx := context.Background()

	cacheKey := "my-key"
	cacheValue := "my-cache-value"

	client := NewMockValkeyClientInterface(ctrl)
	client.EXPECT().Set(ctx, cacheKey, cacheValue, time.Duration(0)).Return(&valkeycompat.StatusCmd{})
	client.EXPECT().SAdd(ctx, "gocache_tag_tag1", "my-key").Return(&valkeycompat.IntCmd{})
	client.EXPECT().Expire(ctx, "gocache_tag_tag1", 720*time.Hour).Return(&valkeycompat.BoolCmd{})

	store := NewValkey(client)

	// When
	err := store.Set(ctx, cacheKey, cacheValue, lib_store.WithTags([]string{"tag1"}))

	// Then
	assert.Nil(t, err)
}

func TestValkeyDelete(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)

	ctx := context.Background()

	cacheKey := "my-key"

	client := NewMockValkeyClientInterface(ctrl)
	client.EXPECT().Del(ctx, "my-key").Return(&valkeycompat.IntCmd{})

	store := NewValkey(client)

	// When
	err := store.Delete(ctx, cacheKey)

	// Then
	assert.Nil(t, err)
}

func TestValkeyInvalidate(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)

	ctx := context.Background()

	cacheKeys := &valkeycompat.StringSliceCmd{}

	client := NewMockValkeyClientInterface(ctrl)
	client.EXPECT().SMembers(ctx, "gocache_tag_tag1").Return(cacheKeys)
	client.EXPECT().Del(ctx, "gocache_tag_tag1").Return(&valkeycompat.IntCmd{})

	store := NewValkey(client)

	// When
	err := store.Invalidate(ctx, lib_store.WithInvalidateTags([]string{"tag1"}))

	// Then
	assert.Nil(t, err)
}

func TestValkeyClear(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	client := NewMockValkeyClientInterface(ctrl)
	store := NewValkey(client)

	tests := []struct {
		name          string
		returnValue   *valkeycompat.StatusCmd
		returnError   error
		expectError   bool
		expectedError string
	}{
		{
			name:        "Successfully clears data",
			returnValue: &valkeycompat.StatusCmd{},
			returnError: nil,
			expectError: false,
		},
		{
			name:          "Returns error on failure",
			returnValue:   NewStatusResult("", fmt.Errorf("flush error")),
			returnError:   nil,
			expectError:   true,
			expectedError: "flush error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.expectError {
				client.EXPECT().FlushAll(ctx).Return(tt.returnValue).Times(1)
			} else {
				client.EXPECT().FlushAll(ctx).Return(tt.returnValue).Times(1)
			}

			err := store.Clear(ctx)

			if tt.expectError {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError, err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestValkeyGetType(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)

	client := NewMockValkeyClientInterface(ctrl)

	store := NewValkey(client)

	// When - Then
	assert.Equal(t, ValkeyType, store.GetType())
}

func TestValkeyGetWithTTL(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	client := NewMockValkeyClientInterface(ctrl)
	store := NewValkey(client)

	t.Run("Returns Value and TTL", func(t *testing.T) {
		testReturnsValueAndTTL(t, ctx, client, store)
	})

	t.Run("Key Not Found", func(t *testing.T) {
		testKeyNotFound(t, ctx, client, store)
	})

	t.Run("Get Error", func(t *testing.T) {
		testGetError(t, ctx, client, store)
	})

	t.Run("TTL Fetch Error", func(t *testing.T) {
		testTTLFetchError(t, ctx, client, store)
	})
}

func testReturnsValueAndTTL(t *testing.T, ctx context.Context, client *MockValkeyClientInterface, store *ValkeyStore) {
	// Given
	key := "my-key"
	returnValue := "value"
	returnTTL := 10 * time.Second
	client.EXPECT().
		Get(ctx, key).
		Return(NewStringResult(returnValue, nil))
	client.EXPECT().
		TTL(ctx, key).
		Return(NewDurationResult(returnTTL, nil))

	// When
	value, ttl, err := store.GetWithTTL(ctx, key)

	// Then
	assert.NoError(t, err)
	assert.Equal(t, returnValue, value)
	assert.Equal(t, returnTTL, ttl)
}

func testKeyNotFound(t *testing.T, ctx context.Context, client *MockValkeyClientInterface, store *ValkeyStore) {
	// Given
	key := "non-existent-key"
	client.EXPECT().
		Get(ctx, key).
		Return(NewStringResult("", valkeycompat.Nil))

	// When
	value, ttl, err := store.GetWithTTL(ctx, key)

	// Then
	assert.Error(t, err)
	assert.Nil(t, value)
	assert.Equal(t, 0*time.Second, ttl)
}

func testGetError(t *testing.T, ctx context.Context, client *MockValkeyClientInterface, store *ValkeyStore) {
	// Given
	key := "my-key"
	client.EXPECT().
		Get(ctx, key).
		Return(NewStringResult("", fmt.Errorf("some error")))

	// When
	value, ttl, err := store.GetWithTTL(ctx, key)

	// Then
	assert.Error(t, err)
	assert.Equal(t, nil, value)
	assert.Equal(t, 0*time.Second, ttl)
}

func testTTLFetchError(t *testing.T, ctx context.Context, client *MockValkeyClientInterface, store *ValkeyStore) {
	// Given
	key := "my-key"
	client.EXPECT().
		Get(ctx, key).
		Return(NewStringResult("", nil))
	client.EXPECT().
		TTL(ctx, key).
		Return(NewDurationResult(0, fmt.Errorf("ttl error")))

	// When
	value, ttl, err := store.GetWithTTL(ctx, key)

	// Then
	assert.Error(t, err)
	assert.Equal(t, nil, value)
	assert.Equal(t, 0*time.Second, ttl)
}
