package main

import (
	"os"

	"github.com/ali-meh/LineBoxClient/internall/client"
)


func main() {

	arguments := os.Args
	addr := "localhost:1898"
	clientName := "aliMiniMaxAlgo"
	if len(arguments) == 2 {
		clientName = arguments[1]
	}
	if len(arguments) == 3 {
		clientName = arguments[1]
		addr = arguments[2]
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
