package main

import (
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
	return 0
}
