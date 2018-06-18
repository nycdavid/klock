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
			return 1
		}
		err = handle(conn)
		if err != nil {
			log.Fatal(err)
			return 1
		}
	}

	return 0
}

func handle(conn net.Conn) error {
	defer conn.Close()
	_, err := conn.Write([]byte("Foobar"))
	if err != nil {
		return err
	}
	return nil
}
