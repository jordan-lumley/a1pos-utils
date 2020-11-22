package monitor

import (
	"testing"
)

func TestMonitors(t *testing.T) {
	testCases := []Monitor{
		Memory{},
		Processor{},
	}

	for _, tc := range testCases {
		go Start(tc)
	}
}
