package samples

import (
	"fmt"
	"time"
)

func printCount(c chan int) {
	num := 0
	for num >= 0 {
		num = <-c
		fmt.Print(num, " ")
	}
}

func ChannelPrint() {
	c := make(chan int)
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, -1}

	go printCount(c)

	for i, v := range a {
		fmt.Printf("\ni:%v\n", i)
		c <- v
	}
	time.Sleep(time.Millisecond * 1)
	fmt.Println("End of ChannelPrint")
}
