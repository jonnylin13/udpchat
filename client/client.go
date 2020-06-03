package client

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"strings"

	"github.com/jonnylin13/udpchat/protocol"
)

func handleError(err error) bool {
	if err != nil {
		fmt.Println(err)
		return true
	}
	return false
}

func moveCursor(screen *bytes.Buffer, x int, y int) {
	fmt.Fprintf(screen, "\033[%d;%dH", x, y)
}

func clearTerminal(output *bufio.Writer) {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Start the client.
func Start(addressStr string) {
	addr, err := net.ResolveUDPAddr("udp4", addressStr)
	pc, err := net.DialUDP("udp4", nil, addr)

	if handleError(err) {
		return
	}

	fmt.Printf("Connected to server: %s\n", pc.RemoteAddr().String())
	defer pc.Close()

	name := ""
	messages := make(chan string, 3)
	cmdText := make(chan string)
	sigint := make(chan os.Signal)
	signal.Notify(sigint, os.Interrupt)

	reader := bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
	var screen *bytes.Buffer = new(bytes.Buffer)

	for {

		go parseCommands(reader, writer, screen, cmdText)
		go read(pc, messages, screen)

		select {
		case signal := <-sigint:
			if name != "" && signal == os.Interrupt {
				pc.Write(protocol.PacketLeave(name).Pack())
			}
			fmt.Printf("Done\n")
			return
		case msg := <-messages:
			fmt.Println(msg)
		case text := <-cmdText:
			data := []byte{}
			texts := strings.SplitAfterN(text, " ", 2)
			cmd := strings.TrimSpace(texts[0])
			if len(texts) > 1 {
				text = strings.TrimSpace(texts[1])
				// Parse the payload
				switch cmd {
				case "message":
					if len(name) == 0 {
						fmt.Printf("Use handshake <name> first\n")
						continue
					}
					// fmt.Println("Sending message...")
					data = protocol.PacketMessage(name, text).Pack()
					break
				case "handshake":
					data = protocol.PacketHandshake(text).Pack()
					name = text
					fmt.Printf("Sending handshake\n")
					break
				default:
					fmt.Printf("Invalid command %s\n", cmd)
					continue
				}
			} else {
				switch cmd {
				case "quit":
					if len(name) > 0 {
						pc.Write(protocol.PacketLeave(name).Pack())
					}
					fmt.Printf("Done\n")
					return
				case "leave":
					data = protocol.PacketLeave(name).Pack()
					name = ""
					break
				default:
					fmt.Printf("Invalid command %s\n", cmd)
					continue
				}
			}

			_, err := pc.Write(data)

			if handleError(err) {
				return
			}
		}
	}

}

func read(pc *net.UDPConn, messages chan string, screen *bytes.Buffer) {
	buf := make([]byte, 1024)
	_, _, err := pc.ReadFrom(buf)

	if handleError(err) {
		return
	}

	switch buf[0] {
	case protocol.Opcodes()["handshake_ack"]:
		fmt.Printf("Connected to lobby\n")
		// resType = "handshake_ack"
		break
	case protocol.Opcodes()["message_ack"]:
		// resType = "message_ack"
		break
	case protocol.Opcodes()["message"]:
		// resType = "message"
		_, end := protocol.UnpackString(buf, 1)
		msg, _ := protocol.UnpackString(buf, end)
		messages <- msg
		break
	case protocol.Opcodes()["leave_ack"]:
		fmt.Printf("Left the lobby\n")
		break
	default:
		break
	}
}

func parseCommands(reader *bufio.Reader, writer *bufio.Writer, screen *bytes.Buffer, cmdText chan string) {

	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	cmdText <- text

}
