package goRoCache

import "testing"

func TestSignalSendsSignalToStopChannel(t *testing.T) {
	c := cacheChannel{
		stopChannel: make(chan bool),
	}
	c.signal(nil)
	select {
	case <-c.stopChannel:
	default:
		t.Errorf("No signal received on the stop channel")
	}
}
func TestSignalStopChannelIsNil(t *testing.T) {
	c := cacheChannel{
		stopChannel: nil,
	}
	c.signal(nil)
}
func TestNewMapCacheEmptyFields(t *testing.T) {
	cache := NewMapCache()

	if cache.cacheMap == nil {
		t.Errorf("cacheMap field is nil")
	}

	if cache.removeChannels == nil {
		t.Errorf("removeChannels field is nil")
	}

	if cache.updateChannels == nil {
		t.Errorf("updateChannels field is nil")
	}
}
func TestNewMapCacheNonNilFields(t *testing.T) {
	cache := NewMapCache()

	if cache.cacheMap == nil {
		t.Errorf("cacheMap field is nil")
	}

	if cache.removeChannels == nil {
		t.Errorf("removeChannels field is nil")
	}

	if cache.updateChannels == nil {
		t.Errorf("updateChannels field is nil")
	}
}
func TestNewErrorWithNilNestedError(t *testing.T) {
	errType := errorType("test")
	msg := "test message"
	result := newError(errType, msg)

	if result.nestedError != nil {
		t.Errorf("Expected nestedError to be nil, but got %v", result.nestedError)
	}

	if result.msg != msg {
		t.Errorf("Expected msg to be %s, but got %s", msg, result.msg)
	}

	if result.errType != errType {
		t.Errorf("Expected errType to be %s, but got %s", errType, result.errType)
	}
}
