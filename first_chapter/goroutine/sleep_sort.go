package Goroutine

import (
	"fmt"
	"time"
)

func main()  {
	s := []string{"h","e","l","l","o","\n"}
	for _, str := range s {
		go func (str string) {
			time.Sleep(time.Duration(str) * time.Second)
			fmt.Println(str)
		}(str)
	}
	time.Sleep(11 * time.Second)
}