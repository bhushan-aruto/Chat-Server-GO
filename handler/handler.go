package handler

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/bhushan-aruto/chat-server-go/service"
)

var mu sync.Mutex

func HandalClient(conn net.Conn) {
	defer conn.Close()
	conn.Write([]byte("WEL COME TO CHAT 🏡 ...💬 \n\n"))
	reader := bufio.NewReader(conn)

	conn.Write([]byte("ENTER YOUR NAME 👤:"))
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	newName := strings.TrimSpace(name)
	service.Clients[conn] = newName
	service.Messages <- fmt.Sprintf("%s has joined the chat.. 💬\n ", newName)
	time.Sleep(time.Second * 1)
	conn.Write([]byte("===================================== \n"))

	conn.Write([]byte("📢 You can start chatting now...💬\n"))
	conn.Write([]byte("Press 1️⃣ anytime to see 👥 online users.\n"))

	conn.Write([]byte("===================================== \n"))

	for {
		msg, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			delete(service.Clients, conn)
			service.Messages <- fmt.Sprintf("%s has left the chat.. 😢💬 \n", newName)
			conn.Write([]byte("===================================== \n"))

			break
		}

		input := strings.TrimSpace(msg)

		if input == "1" {

			FindUCleint(conn)
		} else {

			name := newName + "👤"
			formatedString := fmt.Sprintf("%s :%s", name, msg)
			service.Messages <- formatedString
			conn.Write([]byte("===================================== \n"))
		}

	}

}

func FindUCleint(conn net.Conn) {
	mu.Lock()
	defer mu.Unlock()

	if len(service.Clients) > 1 {
		conn.Write([]byte("👥 Online users:\n"))
		for c, user := range service.Clients {
			if c != conn {
				conn.Write([]byte(fmt.Sprintf("🔸 %v\n", user)))
			}
		}
		conn.Write([]byte("=====================================\n"))
	} else {
		conn.Write([]byte("🙁 You're the only one here. Invite friends!\n"))
		conn.Write([]byte("Send this link to friends 👉 nc 0.0.0.1 8080\n"))
	}
}
