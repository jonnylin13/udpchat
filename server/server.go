package server

import (
	"fmt"
	"log"
	"net"

	"github.com/jonnylin13/udpchat/protocol"
)

func emit(pc net.PacketConn, users []User, data []byte) {
	for _, user := range users {
		pc.WriteTo(data, user.addr)
	}
}

// Start the server.
func Start(port string) {
	pc, err := net.ListenPacket("udp", port)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Listening on port %s...\n", port)
	defer pc.Close()

	messages := make(chan string, 3)
	users := make(chan User, 3)
	leavers := make(chan string, 3)

	userList := []User{}

	for {
		buf := make([]byte, 1024)

		n, addr, err := pc.ReadFrom(buf)

		if err != nil {
			fmt.Println(err)
			continue
		}

		go read(pc, addr, buf[:n], messages, users, leavers, userList)

		select {
		case msg := <-messages:
			fmt.Printf("Message: %s\n", msg)
		case user := <-users:
			fmt.Printf("Handshake received from %s\n", user.name)
			userList = append(userList, user)
			fmt.Printf("Users connected %s\n", GetNames(userList))
		case leaver := <-leavers:
			userList = RemoveUser(userList, leaver)
			fmt.Printf("%s has been removed...\n", leaver)
			fmt.Printf("Users connected %s\n", GetNames(userList))
		}

	}
}

func read(pc net.PacketConn, addr net.Addr, buf []byte, messages chan string, users chan User, leavers chan string, userList []User) {
	// typeStr := "unknown type"
	response := make([]byte, 1024)

	switch buf[0] {
	case protocol.Opcodes()["handshake"]:
		// typeStr = "handshake"
		name, _ := protocol.UnpackString(buf, 1)
		users <- User{name, addr}
		response = protocol.NewPacketHandshakeAck().Pack()
		break
	case protocol.Opcodes()["message"]:
		// typeStr = "message"
		name, end := protocol.UnpackString(buf, 1)
		msg, _ := protocol.UnpackString(buf, end)
		msg = string("<" + name + "> " + msg)
		messages <- msg
		emit(pc, userList, protocol.NewPacketMessage(name, msg).Pack())
		response = protocol.NewPacketMessageAck().Pack()
		break
	case protocol.Opcodes()["leave"]:
		name, _ := protocol.UnpackString(buf, 1)
		leavers <- name
		response = protocol.NewPacketLeaveAck().Pack()
		break
	default:
		response = protocol.NewPacketUnknownRequestAck().Pack()
		break
	}

	// fmt.Printf("Received a request of type: %s\n", typeStr)
	pc.WriteTo(response, addr)
}
