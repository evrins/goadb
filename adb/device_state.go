package adb

import (
	"log"

	"github.com/evrins/goadb/internal/errors"
)

// DeviceState represents one of the 3 possible states adb will report devices.
// A device can be communicated with when it's in StateOnline.
// A USB device will make the following state transitions:
//
//	Plugged in: StateDisconnected->StateOffline->StateOnline
//	Unplugged:  StateOnline->StateDisconnected
//
//go:generate stringer -type=DeviceState
type DeviceState int8

const (
	StateInvalid DeviceState = iota
	StateDisconnected
	StateOffline
	StateOnline
	StateRecovery
	StatUnauthorized
)

var deviceStateStrings = map[string]DeviceState{
	"":             StateDisconnected,
	"offline":      StateOffline,
	"device":       StateOnline,
	"recovery":     StateRecovery,
	"unauthorized": StatUnauthorized,
}

func parseDeviceState(str string) (DeviceState, error) {
	state, ok := deviceStateStrings[str]
	if !ok {
		return StateInvalid, errors.Errorf(errors.ParseError, "invalid device state: %q", state)
	}
	log.Println(str)
	return state, nil
}
