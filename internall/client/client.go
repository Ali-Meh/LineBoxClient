package client

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"net"
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
		c.SendCord(6, 3)
		/*
			c.SendCord(6, 3)
			2-1
			0-0
			 012345678
			0@-@-@-@-@
			1-#-#-#-#-
			2@-@-@-@-@
			3-#-#-#A#-
			4@-@-@-@-@
			5-#-#-#-#-
			6@-@-@-@-@
			7-#-#-#-#-
			8@-@-@-@-@
		*/
	}
}

//SendCord sends the cordinates ai selected to ai
func (c *Client) SendCord(x, y int8) {
	coords := fmt.Sprintf("%d-%d", y, x)
	encMsg := base64.StdEncoding.EncodeToString([]byte(coords)) + "\n"
	fmt.Fprintf(c.c, encMsg)
}