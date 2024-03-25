package goRoCache

import "testing"

func TestNewRedisCache_NilClient(t *testing.T) {
	address := ""
	password := "password"
	db := 0

	cache := NewRedisCache(address, password, db)

	if cache.client != nil {
		t.Errorf("Expected client to be nil, got %v", cache.client)
	}
}
func TestNewRedisCache_NegativeDB(t *testing.T) {
	address := "localhost:6379"
	password := "password"
	db := -1

	cache := NewRedisCache(address, password, db)

	if cache.client != nil {
		t.Errorf("Expected client to be nil, got %v", cache.client)
	}
}
func TestNewRedisCache_InitializedClient(t *testing.T) {
	address := "localhost:6379"
	password := "password"
	db := 0

	cache := NewRedisCache(address, password, db)

	if cache.client == nil {
		t.Errorf("Expected client to be initialized, got nil")
	}

	options := cache.client.Options()
	if options.Addr != address {
		t.Errorf("Expected client address to be %s, got %s", address, options.Addr)
	}

	if options.Password != password {
		t.Errorf("Expected client password to be %s, got %s", password, options.Password)
	}

	if options.DB != db {
		t.Errorf("Expected client db to be %d, got %d", db, options.DB)
	}
}
func TestNewRedisCache_EmptyMaps(t *testing.T) {
	cache := NewRedisCache("localhost:6379", "password", 0)

	if len(cache.keysSet) != 0 {
		t.Errorf("Expected keysSet to be empty, got %v", cache.keysSet)
	}

	if len(cache.removeChannels) != 0 {
		t.Errorf("Expected removeChannels to be empty, got %v", cache.removeChannels)
	}
}
