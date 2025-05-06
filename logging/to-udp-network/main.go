package main

import (
	"log"
	"net"
	"time"
)

// start server to listen in terminal:
// $ nc -luk 1902

func main() {
	timeout := 30 * time.Second
	conn, err := net.DialTimeout("udp", "localhost:1902", timeout)
	if err != nil {
		panic("Failed to connect to localhost:1902")
	}
	defer conn.Close()

	f := log.Ldate | log.Lshortfile
	logger := log.New(conn, "test-prefix ", f)

	logger.Println("Regular message")
	logger.Panicln("Oops! Panic.")
}
