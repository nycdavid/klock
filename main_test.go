package main

import (
	"fmt"
	"testing"
)

func TestMain_requiresPort(t *testing.T) {
	expected := 1
	got := realMain()

	if expected != got {
		msg := fmt.Sprintf("Expected exit code %d, got %d", expected, got)
		t.Error(msg)
	}
}