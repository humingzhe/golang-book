package main

import (
	"fmt"
	"time"
)
/*固定时长定时器第一种写法*/
func main() {
	// 创建3秒钟定时器
	timer := time.NewTimer(3 * time.Second)
	fmt.Println("定时器创建完毕")
	// 打印当前时间
	fmt.Println(time.Now())

	// 阻塞3秒后才能读出当前时间
	x := <-timer.C
	fmt.Println(x)
}

/*固定时长定时器第二种写法*/

func Fixed() {
	fmt.Println(time.Now())
	//y := <-time.NewTimer(3 * time.Second).C
	// 3秒钟后读出时间
	x := <-time.After(3 * time.Second)
	fmt.Println(x)
	//fmt.Println(y)

}