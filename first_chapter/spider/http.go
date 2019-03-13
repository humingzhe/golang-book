package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	// 声明一个Handle，默认是根目录。把根目录请求转换到Handle上，然后直接用http.FileServer，FileServer指向当前目录。
	http.Handle("/", http.FileServer(http.Dir(".")))
	// 打印log.Fatal，Fatal返回http.ListenAndServe。os.Args是地址，地址通过第一个参数传进来。
	log.Fatal(http.ListenAndServe(os.Args[1], nil))
}
