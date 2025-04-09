package concurrency

import (
	"fmt"
	"runtime"
)

func AnonGoRoutine() {
	fmt.Println("Outside before goroutine.")

	go func() {
		fmt.Println("Inside goroutine")
	}()
	fmt.Println("Outside after goroutine")
	runtime.Gosched()
}
