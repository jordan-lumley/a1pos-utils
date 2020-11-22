package logger

import "testing"

func TestInstance(t *testing.T) {
	Instance().Info("Successful log test")
}
