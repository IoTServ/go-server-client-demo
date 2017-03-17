package main

import (
	"fmt"
	"net"
	"os"
	"runtime"
)

type Channel struct {
	from, to net.Conn
}

func pass_through(c *Channel) {
	b := make([]byte, 10240)
	for {
		n, err := c.from.Read(b)
		if err != nil {
			break
		}
		if n > 0 {
			c.to.Write(b[:n])

		}
	}
	c.from.Close()
	c.to.Close()
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	c, err := net.Listen("tcp", ":2000")
	if err != nil {
		fmt.Printf("Unable to start listener, %v\n", err)
		os.Exit(1)
	}
	u, err := net.Listen("tcp", ":3000")
	if err != nil {
		fmt.Printf("Unable to start listener, %v\n", err)
		os.Exit(1)
	}
	for {
		client, err := c.Accept()
		if err != nil {
			fmt.Printf("client conn failed, %v\n", err)
		}
		user, err := u.Accept()
		if err != nil {
			fmt.Printf("user conn failed, %v\n", err)
		}

		go pass_through(&Channel{user, client})
		go pass_through(&Channel{client, user})

	}
}
