package goRoCache

import "testing"

func TestNewRedisCache_InitializedFields(t *testing.T) {
	address := "localhost:6379"
	password := "password"
	db := 0

	cache := NewRedisCache(address, password, db)

	if cache.keysSet == nil {
		t.Errorf("Expected keysSet to be initialized, got nil")
	}

	if cache.removeChannels == nil {
		t.Errorf("Expected removeChannels to be initialized, got nil")
	}

	if cache.client == nil {
		t.Errorf("Expected client to be initialized, got nil")
	}
}
