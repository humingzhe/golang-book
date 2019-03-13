package main

import (
	"bufio"
	"log"
	"net/rpc"
	"os"
	"fmt"
)
func main() {
	// 首先RPC这个地方需要先进行拨号，这里用的系统自带RPC框架。
	// 先去连接，先告诉你我是TCP，我要连接localhost的42586端口。
	client, err := rpc.Dial("tcp", "localhost:42586")
	// 处理错误
	if err != nil {
		log.Fatal(err)
	}
	// 从标准输入里面读内容
	in := bufio.NewReader(os.Stdin)
	for {
		// 读完进行ReadLine，ReadLine表示一行行读。
		line, _, err := in.ReadLine()
		if err != nil {
			log.Fatal(err)
		}
		var reply string
		// RPC Client在这里会调用，因为这里有.Call，在这里面有一个Listener.GerLine
		// 再把Line扔回去，最后用reply传一个引用过来，reply本身是一个string类型。
		// 这么一调用，相当于是error就会返回reply调用
		// 就相当于是你要在本地就是叫Listener.GetLine(Line)，reply等于它
		err = client.Call("Listener.GetLine", line, &reply)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(reply)
	}
}
