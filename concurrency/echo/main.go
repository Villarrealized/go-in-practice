package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	done := time.After(30 * time.Second)
	echo := make(chan []byte)
	go readStdIn(echo)
	for {
		select {
		case buf := <-echo:
			os.Stdout.Write(buf)
		case <-done:
			fmt.Println("Timed out")
			os.Exit(0)
		}
	}
}

func echoInOut(in io.Reader, out io.Writer) {
	io.Copy(out, in)
}

func echoStdIn() {
	go echoInOut(os.Stdin, os.Stdout)
	time.Sleep(30 * time.Second)
	fmt.Println("Timed out.")
	os.Exit(0)
}

func readStdIn(out chan<- []byte) {
	for {
		data := make([]byte, 1024)
		l, _ := os.Stdin.Read(data)
		if l > 0 {
			out <- data
		}
	}
}
