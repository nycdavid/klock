package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	os.Exit(realMain())
}

func realMain() int {
	port := os.Getenv("PORT")
	if port == "" {
		return 1
	}
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", port))
	if err != nil {
		log.Fatal(err)
		return 1
	}
	_, err = listener.Accept()
	if err != nil {
		log.Fatal(err)
		return 1
	}
	return 0
}
