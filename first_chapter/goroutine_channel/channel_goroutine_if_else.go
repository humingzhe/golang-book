package Goroutine_Channel

import (
	"fmt"
	"time"
)

func main() {
	s := "Channel\n"
	fmt.Printf("%T %v\n", s[0], s[0])
	first := make(chan int)
	in := first
	for i, n := range s{
		var out chan int
		if i == len(s)-1 {
			out = first
		} else {
			out = make(chan int)
		}
		go func(i int, n rune, in chan int, out chan int) {
			for {
				in_num := <-in
				fmt.Printf("%v read from in %v\n", string(n), in_num)
				fmt.Println(string(n))
				time.Sleep(time.Second)
				fmt.Printf("%v write to next %v\n", string(n), i)
				out <- i
			}
		}(i, n, in, out)
		in = out
	}
	fmt.Printf("%v write to begin\n", -1)
	first <- -1
	time.Sleep(10 * time.Second)
	fmt.Println(s)
}
