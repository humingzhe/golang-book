package Goroutine_Channel

import (
	"fmt"
	"time"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	fmt.Printf("start of %v\n", s)
	c <- sum // sed sum to c
	fmt.Printf("end of %v\n", s)
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6}

	c1 := make(chan int)
	c2 := make(chan int)
	go sum(s[:len(s)/2], c1)
	go sum(s[len(s)/2:], c2)
	time.Sleep(time.Second)
	x, y := <-c1, <-c2 //receive from c

	fmt.Println(x, y, x+y)
	time.Sleep(time.Second)
}
