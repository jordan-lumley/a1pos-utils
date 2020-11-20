package monitor

import "testing"

func TestExecute(t *testing.T) {
	err := Execute()
	if err != nil {
		t.Error("Failed to execute system monitor")
	}
}
