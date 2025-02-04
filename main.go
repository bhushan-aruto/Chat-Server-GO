package main

import (
	"fmt"
	"log"
	"net"

	"github.com/bhushan-aruto/chat-server-go/handler"
	"github.com/bhushan-aruto/chat-server-go/service"
)

// var wg sync.WaitGroup

func main() {

	listener, err := net.Listen("tcp", "0.0.0.0:8080")

	if err != nil {
		log.Println(err)
		return
	}
	defer listener.Close()

	// broadcast

	fmt.Println(" 💬 server started running ..")
	go service.Broadcast()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("connection failed :", err)
			continue
		}
		go handler.HandalClient(conn)

	}

}
