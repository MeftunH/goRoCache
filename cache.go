package goRoCache

import "time"

type Cache interface {
	Store(key, val interface{}) error
	Get(key interface{}) (interface{}, error)
	Remove(key interface{}) error
	Replace(key, val interface{}) error
	Clear() error
	Keys() ([]interface{}, error)
	StoreWithExpiration(key interface{}, item lfuItem, ttl time.Duration) interface{}
}

type CacheStoreWithExpiration interface {
	StoreWithExpiration(key, val interface{}, ttl time.Duration) error
}

type ExpiringCache interface {
	Cache
	StoreWithExpiration(key, val interface{}, ttl time.Duration) error
	ReplaceWithExpiration(key, val interface{}, ttl time.Duration) error
	Expire(key interface{}, ttl time.Duration) error
}

type UpdatingCache interface {
	Cache
	StoreWithUpdate(key, initialValue interface{},
		updateFunc func(currValue interface{}) interface{},
		period time.Duration) error
	ReplaceWithUpdate(key, initialValue interface{},
		updateFunc func(currValue interface{}) interface{},
		period time.Duration) error
}
type errorType string

const (
	errorTypeUnexpectedError   errorType = "UnexpectedError"
	errorTypeAlreadyExists               = "AlreadyExists"
	errorTypeDoesNotExist                = "DoesNotExist"
	errorTypeNonPositivePeriod           = "NonPositivePeriod"
	errorTypeNilUpdateFunc               = "NilUpdateFunc"
	errorTypeInvalidKeyType              = "InvalidKeyType"
	errorTypeInvalidMessage              = "InvalidMessage"
	errorTypeCacheNotEmpty               = "CacheNotEmpty"
)

type cacheError struct {
	msg         string
	errType     errorType
	nestedError error
}

func (ce cacheError) Error() string {
	return ce.msg
}
func newError(errType errorType, msg string) cacheError {
	return cacheError{
		msg:     msg,
		errType: errType,
	}
}
func newWrapperError(errType errorType, msg string, nestedError error) cacheError {
	return cacheError{
		msg:         msg,
		errType:     errType,
		nestedError: nestedError,
	}
}
