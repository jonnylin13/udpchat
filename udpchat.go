package main

import (
	"fmt"
	"os"

	"github.com/jonnylin13/udpchat/client"
	"github.com/jonnylin13/udpchat/server"
)

func main() {
	arguments := os.Args

	if len(arguments) != 3 {
		fmt.Println("udpchat <client/server> <host/port>")
		return
	}

	option := arguments[1]
	addrStr := arguments[2]

	if option == "server" {
		server.Start(addrStr)
	} else if option == "client" {
		client.Start(addrStr)
	} else {
		fmt.Println("Invalid option.")
	}
}
