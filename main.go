package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	os.Exit(realMain())
}

func realMain() int {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT required")
		return 1
	}
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", port))
	if err != nil {
		log.Fatal(err)
		return 1
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		go handle(conn)
		log.Print("After")
	}

	return 0
}

func handle(conn net.Conn) {
	defer conn.Close()
	for {
		_, err := io.WriteString(conn, "Foobar\n")
		if err != nil {
			log.Print("Disconnect")
			return
		}
		time.Sleep(2 * time.Second)
	}
}
