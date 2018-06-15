package main

import (
	"fmt"
	"log"
	"net"
	"os"
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

func TestMain_serverWritesFoobar(t *testing.T) {
	os.Setenv("PORT", "3000")
	go func() {
		for {
			conn, err := net.Dial("tcp", "localhost:3000")
			if err != nil {
				log.Fatal(err)
				continue
			}
			fmt.Println(conn)
		}
	}()
	realMain()
}
