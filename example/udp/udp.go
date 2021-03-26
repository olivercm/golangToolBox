package main

import (
	"log"
	"net"
)

// 终端执行：echo -n "haha" | nc -u -w1 127.0.0.1 20000
func main() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{Port: 20000})
	if err != nil {
		log.Fatalf("listen error: %v\n", err)
	}
	defer listen.Close()

	for {
		var buf [1024]byte
		n, addr, err := listen.ReadFromUDP(buf[:])
		if err != nil {
			log.Printf("read udp error: %v\n", err)
			continue
		}
		data := append([]byte("hello "), buf[:n]...)
		listen.WriteToUDP(data, addr)
	}
}
