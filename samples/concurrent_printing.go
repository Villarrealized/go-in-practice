package samples

import (
	"fmt"
	"time"
)

func count() {
	for i := range 5 {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 1)
	}
}

func PrintConcurrent() {
	go count()
	time.Sleep(time.Millisecond * 2)
	fmt.Println("Hello World")
	time.Sleep(time.Millisecond * 5)
}
