package Goroutine_Channel

import "fmt"

func fibonnaci(c, quit chan int) { // C和Quit都是Channel int
	x, y:= 0, 1
	for {
		select { // 在Select里关注两个Channel
		case c <- x: // 首先关注C，把X往C里面塞
			x, y = y, x+y // 如果X往C里面写成功了，Select就会走这个分支，这个Select就相当于是个循环
		// 从quit这个channel里去读东西
		case <- quit: // 什么都不要就直接扔掉了
			fmt.Println("quit") // 收到Quit Channel如果有消息的话就会打印Print
			return //打印完Print后Return出去
			// 到此为止fibonaci这个函数就退出了
		}
	}
}

func main() {
	c := make(chan int) // 用C声明一个Channel
	quit := make(chan int) // 用Quit声明一个Channel
	go func() { // 执行Go func
		for i := 0; i < 10; i++ { // function里不断去写，循环10次
			fmt.Println(<-c) // 从Channel里读东西
		}
		quit <- 0 // 读完之后往Quit里写0
	}()
	fibonnaci(c,quit) // 写完0后把斐波那契直接当成主协程在这运行
}
