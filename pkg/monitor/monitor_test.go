package monitor

import (
	"testing"
)

func TestMonitors(t *testing.T) {
	testCases := []Monitor{
		Memory{Output: make(chan float64)},
		Processor{Output: make(chan float64)},
	}

	for _, tc := range testCases {
		go Start(tc)
	}
}
