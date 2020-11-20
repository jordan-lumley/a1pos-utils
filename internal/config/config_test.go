package config

import "testing"

func TestConfig(t *testing.T) {
	err := Config()
	if err != nil {
		t.Error(err)
	}
}
