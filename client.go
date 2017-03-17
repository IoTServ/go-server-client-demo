package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"runtime"
)

var host *string = flag.String("host", "127.0.0.1", "target host or address")

func pass_through(server net.Conn, browser chan net.Conn, ask chan bool) {
	retur := true
	b := make([]byte, 10240)
	var brow net.Conn
	for {
		n, err := server.Read(b)

		if retur {

			brow, err = net.Dial("tcp", "127.0.0.1:80")
			if err != nil {
				fmt.Printf("Unable to start dial, %v\n", err)
				os.Exit(1)
			}
			browser <- brow
			ask <- true
		}
		retur = false
		if err != nil {
			break
		}
		if n > 0 {
			brow.Write(b[:n])

		}
	}
	fmt.Println("远程服务器其中一条tcp以关闭，关闭此条tcp链接")
	server.Close()
	brow.Close()

}
func pass_through1(server net.Conn, browser chan net.Conn) {

	b := make([]byte, 10240)
	brow := <-browser
	for {
		n, err := brow.Read(b)
		if err != nil {
			break
		}
		if n > 0 {
			server.Write(b[:n])
		}

	}
	fmt.Println("本地网站其中一条tcp以关闭，关闭此条tcp链接")
	server.Close()
	brow.Close()
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	browser := make(chan net.Conn)
	ask := make(chan bool)
	flag.Parse()
	if flag.NFlag() != 1 {
		fmt.Printf("usage: -host 远程服务器ip\n")
		flag.PrintDefaults()
		os.Exit(1)
	}
	target := net.JoinHostPort(*host, "2000")
	for {

		server, err := net.Dial("tcp", target)
		if err != nil {
			fmt.Printf("Unable to dial server, %v\n", err)
			os.Exit(1)
		}

		go pass_through(server, browser, ask)
		go pass_through1(server, browser)
		<-ask
	}
}
