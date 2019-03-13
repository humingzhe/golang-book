package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建一个1秒为周期的秒表
	ticker := time.NewTicker(1 * time.Second)

	var i int
	for {
		// 每隔一秒读出当前时间
		x := <- ticker.C
		//fmt.Println("\n" ,x)
		fmt.Println("\r" ,x)


		// 10秒后停止计时并退出
		i++
		if i> 9{
			// 停掉秒表会导致ticker.C永远无法读出数据，执着要读会导致死锁(deadlock)
			ticker.Stop()
			break
		}
	}
	fmt.Println("\n计时结束")
}