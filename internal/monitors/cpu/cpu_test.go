package cpu

import "testing"

func TestStart(t *testing.T) {
	cpuChan := make(chan float64)

	cpuMon, err := New(cpuChan)
	if err != nil {
		t.Error(err)
	}

	cpuMon.Start()
}
