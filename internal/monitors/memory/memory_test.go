package memory

import "testing"

func TestStart(t *testing.T) {
	memChan := make(chan float64)

	memMon, err := New(memChan)
	if err != nil {
		t.Error(err)
	}

	memMon.Start()
}
