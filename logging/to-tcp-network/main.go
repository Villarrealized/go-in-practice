package main

import (
	"log"
	"net"
)

// start server to listen in terminal:
// $ nc -lk 1902

func main() {
	conn, err := net.Dial("tcp", "localhost:1902")
	if err != nil {
		panic("Failed to connect to localhost:1902")
	}
	defer conn.Close()

	f := log.Ldate | log.Lshortfile
	logger := log.New(conn, "test-prefix ", f)

	logger.Println("Regular message")
	logger.Panicln("Oops! Panic.")
}
