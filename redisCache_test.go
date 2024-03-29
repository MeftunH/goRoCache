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
func TestRedisCache_Store(t *testing.T) {
	cache := NewRedisCache("localhost:6379", "password", 0)

	key := "testKey"
	value := "testValue"

	err := cache.Store(key, value)

	if err != nil {
		t.Errorf("Store returned an error: %v", err)
	}

	storedValue, err := cache.Get(key)
	_ = err

	if err != nil {
		t.Errorf("Get returned an error: %v", err)
	}

	if storedValue != value {
		t.Errorf("Stored value does not match original value. Got %v, want %v", storedValue, value)
	}

}
func TestRedisCache_Get_ExistingKey(t *testing.T) {
	// Create a new RedisCache instance
	cache := NewRedisCache("localhost:6379", "password", 0)

	// Define the key and value to be stored
	key := "testKey"
	value := "testValue"

	// Store the key-value pair
	err := cache.Store(key, value)
	if err != nil {
		t.Fatalf("Store returned an error: %v", err)
	}

	// Retrieve the value from the cache
	storedValue, err := cache.Get(key)
	_ = err

	// Check if there was an error
	if err != nil {
		t.Errorf("Get returned an error: %v", err)
	}

	// Check if the stored value matches the original value
	if storedValue != value {
		t.Errorf("Stored value does not match original value. Got %v, want %v", storedValue, value)
	}
}

func TestRedisCache_Get_NonExistentKey(t *testing.T) {
	// Create a new RedisCache instance
	cache := NewRedisCache("localhost:6379", "password", 0)

	// Define a key that does not exist in the cache
	key := "nonExistentKey"

	// Attempt to retrieve the value from the cache
	_, err := cache.Get(key)

	// Check if there was an error
	if err == nil {
		t.Errorf("Expected Get to return an error for non-existent key, got nil")
	}
}

func TestRedisCache_Get_RemovedKey(t *testing.T) {
	// Create a new RedisCache instance
	cache := NewRedisCache("localhost:6379", "password", 0)

	// Define the key and value to be stored
	key := "testKey"
	value := "testValue"

	// Store the key-value pair
	err := cache.Store(key, value)
	if err != nil {
		t.Fatalf("Store returned an error: %v", err)
	}

	// Remove the key-value pair
	err = cache.Remove(key)
	if err != nil {
		t.Fatalf("Remove returned an error: %v", err)
	}

	// Attempt to retrieve the value from the cache
	_, err = cache.Get(key)

	// Check if there was an error
	if err == nil {
		t.Errorf("Expected Get to return an error for removed key, got nil")
	}
}
