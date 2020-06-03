package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/jonnylin13/udpchat/client"
	"github.com/jonnylin13/udpchat/server"
)

func main() {
	arguments := os.Args

	if len(arguments) != 3 {
		fmt.Println("udpchat <client/server> <host:port>")
		return
	}

	option := arguments[1]
	addrStr := arguments[2]

	if !strings.Contains(addrStr, ":") {
		// Load from file
		jsonFile, err := os.Open("rooms.json")

		if err != nil {
			fmt.Println(err)
			return
		}

		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)
		var result map[string]string
		json.Unmarshal([]byte(byteValue), &result)
		if addr, ok := result[addrStr]; ok {
			addrStr = addr
		}
	}

	switch option {
	case "server":
		server.Start(addrStr)
		break
	case "client":
		client.Start(addrStr)
		break
	default:
		fmt.Println("Invalid command.")
		break
	}
}
