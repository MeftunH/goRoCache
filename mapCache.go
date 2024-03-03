package goRoCache

import (
	_ "fmt"
	_ "time"
)

type cacheChannel struct {
	stopChannel chan bool
}
