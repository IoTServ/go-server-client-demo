package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

var host *string = flag.String("host", "127.0.0.1", "target host or address")

func main() {
	flag.Parse()
	if flag.NFlag() != 1 {
		fmt.Printf("usage: -host 远程服务器ip\n")
		flag.PrintDefaults()
		os.Exit(1)
	}
	target := net.JoinHostPort(*host, "2000")
	server, err := net.Dial("tcp", target)
	if err != nil {
		log.Fatal(err)
	}

	// 显示连接发送来的内容
	for {
		buf := make([]byte, 128000)
		n, err := server.Read(buf)

		if err != nil {
			fmt.Println("server read")
			log.Fatal(err)
		}
		goun(n, buf, server)
	}

}
func goun(n int, buf []byte, server net.Conn) {
	brower, err := net.Dial("tcp", "127.0.0.1:80")
	if err != nil {
		fmt.Println("conn 80")
		fmt.Println(err)
	}
	b := buf[:n]
	_, err = brower.Write(b)
	if err != nil {
		fmt.Println("brower.Write")
		fmt.Println(err)
	}
	n, err = brower.Read(buf)
	if err != nil {
		fmt.Println("brower.Read")
		fmt.Println(err)
	}
	b = buf[:n]
	_, err = server.Write(b)
	if err != nil {
		fmt.Println("server.Write")
		fmt.Println(err)
	}
	// 关闭连接
	brower.Close()
}
