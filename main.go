package main

import (
	"os"

	"github.com/ali-meh/LineBoxClient/client"
)

func main() {
	arguments := os.Args
	addr := "localhost:1898"
	clientName := "aliMiniMaxAlgo"
	if len(arguments) == 2 {
		addr = arguments[1]
	}
	if len(arguments) == 3 {
		addr = arguments[1]
		clientName = arguments[2]
	}

	client, err := client.NewClientAddress(clientName, addr)
	checkError(err)
	client.ReadServer()

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
