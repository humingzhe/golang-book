package Goroutine

import (
	"fmt"
	"time"
)

func say(s string, c int) {
	for i := 0; i < c; i++ {
		fmt.Println(s)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	say("hello", 5)
	go say("world", 10)
	fmt.Println("hello humingzhe")
}