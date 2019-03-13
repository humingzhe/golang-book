package Goroutine

import (
	"fmt"
	"time"
)

func main() {
	s := "Channel\n"
	fmt.Printf("%T\n", s[0])
	for i, n := range s{
		go func(i int, n rune) {
			time.Sleep(time.Duration(i) * time.Millisecond)
			fmt.Println(string(n))
		}(i,n)
	}
	time.Sleep(1 * time.Second)
	fmt.Println(s)
}
