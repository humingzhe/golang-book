package Goroutine_Channel

import "fmt"

func fibonnaci(n int,c chan int) {
	x, y:= 0, 1
	for i := 0; i < n; i++ { // 根据它的容量循环n次
		c <- x // 把斐波那契求出来的值写进去
		x, y = y, x+y // x=y，y=x+y。
	}
	close(c) // 当循环到c := make(chan int, 10) 等于10之后就会进行close。Channel可以再这里进行Close，Close之后，range或者任何读的东西都会立刻返回，什么都读不到。
}

func main() {
	c := make(chan int, 10) // buffer10的Channel
	go fibonnaci(cap(c), c) // go fibonaci，cap表示求Channel容量，也就是这个10。再把Channel传进去
	for i := range c { //死循环去读
		fmt.Println(i)
	}
}
