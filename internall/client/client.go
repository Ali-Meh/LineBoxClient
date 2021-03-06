package client

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"net"

	"github.com/ali-meh/LineBoxClient/internall/gamemap"
	"github.com/ali-meh/LineBoxClient/internall/mcts"
)

//Client keeps track of the server and client name
type Client struct {
	Name  string
	c     net.Conn
	first bool
}

//NewClient declare your Client
func NewClient(name string, c net.Conn) *Client {
	client := new(Client)
	client.Name = name
	client.c = c
	client.nameClient()
	return client
}

//NewClientAddress Create client with address
func NewClientAddress(name, host string) (*Client, error) {
	client := new(Client)
	client.Name = name
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}
	client.c = conn
	client.nameClient()
	return client, nil
}

func (c *Client) nameClient() {
	fmt.Fprintf(c.c, c.Name+"\n")
}

//ReadServer read data from Server
func (c *Client) ReadServer() /*  (string, error)  */ {
	for {
		encMsg, _ := bufio.NewReader(c.c).ReadString('\n')
		message, _ := base64.StdEncoding.DecodeString(encMsg)

		fmt.Printf("server ->: %s\n", message)
		gmap := gamemap.NewMapSquare(4)
		maximizer := "A"
		if message[0] == '2' {
			maximizer = "B"
		}
		gmap.Update(string(message), maximizer)
		// depth := 1.0
		// depth := (float64(len(gmap.AIndexes)+len(gmap.BIndexes))/float64(len(gmap.Cells)*len(gmap.Cells[0])*4))*3 + 3
		// fmt.Println("Depth set to ", depth)
		// move, _ := ai.SelectMove(*gmap, int(depth), string(message[0:2]))

		move := mcts.SelectMove(*gmap, maximizer)
		fmt.Printf("Sending %v to server\n", move)
		c.SendCord(move[0], move[1])
	}
}

//SendCord sends the cordinates ai selected to ai
func (c *Client) SendCord(x, y int8) {
	coords := fmt.Sprintf("%d-%d", y, x)
	encMsg := base64.StdEncoding.EncodeToString([]byte(coords)) + "\n"
	fmt.Fprintf(c.c, encMsg)
}
