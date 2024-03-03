package goRoCache

import (
	_ "fmt"
	"sync"
	_ "time"
)

type cacheChannel struct {
	stopChannel chan bool
}
type mapCache struct {
	// Holds the key/values in the cache
	cacheMap map[interface{}]interface{}

	// Holds the channels that stop the auto removal routines.
	removeChannels map[interface{}]*cacheChannel

	// Holds the channels that stop the auto update routines.
	updateChannels map[interface{}]*cacheChannel

	mutex sync.Mutex
}
