package main

import (
	"log"
	"os"
)

func main() {
	logfile, _ := os.Create("./log.txt")
	defer logfile.Close()

	logger := log.New(logfile, "test-prefix ", log.LstdFlags|log.Lshortfile)

	logger.Println("Regular message")
	logger.Fatalln("Oops! Fatal error.")
}
