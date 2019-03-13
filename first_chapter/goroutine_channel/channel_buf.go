package Goroutine_Channel

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	for c := range ch {
		fmt.Println(c)
	}
}
