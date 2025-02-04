package service

import (
	"fmt"
	"net"
)

var Messages = make(chan string)
var Clients = make(map[net.Conn]string)

func Broadcast() {
	for {
		mssg := <-Messages
		for conn := range Clients {
			fmt.Fprintln(conn, mssg)
		}
	}
}
