package main

// "github.com/ali-meh/LineBoxClient/client"
import (
	"fmt"

	gmap "github.com/ali-meh/LineBoxClient/internall/gamemap"
)

func main() {
	gamemap := gmap.NewMapSquare(3)
	fmt.Println(gamemap)

	// arguments := os.Args
	// addr := "localhost:1898"
	// clientName := "aliMiniMaxAlgo"
	// if len(arguments) == 2 {
	// 	addr = arguments[1]
	// }
	// if len(arguments) == 3 {
	// 	addr = arguments[1]
	// 	clientName = arguments[2]
	// }

	// client, err := client.NewClientAddress(clientName, addr)
	// checkError(err)
	// client.ReadServer()

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
