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
