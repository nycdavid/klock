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

func TestMain_writesErrorToStdout(t *testing.T) {
	t.Skip("Pending")
}

func TestMain_serverWritesFoobar(t *testing.T) {
	t.Skip("Pending")
}
