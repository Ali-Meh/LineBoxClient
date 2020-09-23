package client

import (
	"fmt"
	"io"
	"net"
	"time"
)

func (c *Client) upgradeAlive() error {
	conn := c.c
	err := conn.(*net.TCPConn).SetKeepAlive(true)
	if err != nil {
		return err
	}

	err = conn.(*net.TCPConn).SetKeepAlivePeriod(30 * time.Second)
	if err != nil {
		return err
	}
	notify := make(chan error)

	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := conn.Read(buf)
			if err != nil {
				notify <- err
				if io.EOF == err {
					return
				}
			}
			if n > 0 {
				fmt.Printf("unexpected data: %s\n", buf[:n])
			}
		}
	}()

	for {
		select {
		case err := <-notify:
			fmt.Println("connection dropped message", err)
			break
		case <-time.After(time.Second * 1):
			fmt.Println("timeout 1, still alive")
		}
	}
}
