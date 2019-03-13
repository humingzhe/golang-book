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
		//先make一个Channel
		out := make(chan  int)
		//在返回一个Channel，out通过外部传进来
		go func(i int, n rune, in chan int, out chan int) {
			// 拿到Channel后去读Channel，不关心内容就可以直接往里面扔
			in_num := <- in
			// 以日志的形式展现出来，使之更加清晰易懂整个程序的执行过程。
			fmt.Printf("%v read from in %v\n", string(n), in_num)
			fmt.Println(string(n))
			fmt.Printf("%v write to next %v\n", string(n), i)
			// Print后需要告诉下一个人可以启动了，下个是怎么弄出来的呢？先make一个Channel
			// ret_c := make(chan int)
			// make完之后，往里面塞东西，可以随便塞
			// ret_c <- 0
			out <- i // 塞字母是协程创建时，就知道它要输出什么了。
		}(i, n, in, out) // 增加参数 in 和 out
		// 将in 转化成out
		in = out
	}
	first <- -1
	fmt.Printf("%v write to begin\n", -1)
	time.Sleep(1 * time.Second)

	fmt.Println(s)
}
