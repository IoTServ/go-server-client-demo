package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	c, err := net.Listen("tcp", ":2000")

	if err != nil {
		fmt.Println("listen 2000")
		log.Fatal(err)
	}
	u, err := net.Listen("tcp", ":3000")
	if err != nil {
		fmt.Println("listen 3000")
		log.Fatal(err)
	}
	client, err := c.Accept()
	fmt.Println("client已连接")
	if err != nil {
		log.Fatal(err)
	}

	for {
		user, err := u.Accept()
		fmt.Println("user已连接")
		if err != nil {
			fmt.Println("user连接出错")
			fmt.Println(err)
			continue
		}
		goun(user, client)

	}

}

func goun(user, client net.Conn) {
	buf := make([]byte, 128000)
	// 显示连接发送来的内容
	n, err := user.Read(buf)
	if err != nil {
		fmt.Println("user.Read")
		fmt.Println(err)
		return
	}
	b := buf[:n]
	_, err = client.Write(b)
	if err != nil {
		fmt.Println("client.Write")
		log.Println(err)
	}
	n, err = client.Read(buf)
	if err != nil {
		fmt.Println("client.Read")
		log.Println(err)
	}
	b = buf[:n]
	fmt.Println("收到的数据长度", n)
	_, err = user.Write(b)
	if err != nil {
		fmt.Println("user.Write")
		fmt.Println(err)
		return
	}
	// 关闭连接
	user.Close()
}
