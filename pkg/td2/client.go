package td2

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"strings"
	"time"
)

type Client struct {
	ctx         context.Context
	conn        net.Conn
	stationHash []byte
	ReadChan    chan string
	IsConnected bool
}

func New(ctx context.Context, stationHash string) *Client {
	return &Client{
		ctx:         ctx,
		stationHash: []byte(stationHash),
		ReadChan:    make(chan string),
		IsConnected: false,
	}
}

func (c *Client) Connect(address string) error {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return err
	}

	c.conn = conn
	c.IsConnected = true

	go c.readMessages()

	return nil
}

func (c *Client) Disconnect() {
	c.IsConnected = false
	c.conn.Close()
}

func (c *Client) Write(data string) error {
	c.conn.SetWriteDeadline(time.Now().Add(1 * time.Second))
	_, err := c.conn.Write([]byte(data + "\r\n"))
	return err
}

func (c *Client) readMessages() {
	for {
		scanner := bufio.NewScanner(c.conn)

		for {
			ok := scanner.Scan()
			text := scanner.Text()

			if strings.HasPrefix(text, "/") {
				c.handleCommands(text)
			} else {
				c.ReadChan <- text
			}

			if !ok {
				break
			}
		}
	}
}

func (c *Client) handleCommands(message string) {
	switch {
	case strings.HasPrefix(message, "/SHA2"):
		c.Write(fmt.Sprintf("/SHA2 %s", c.stationHash))
	}
}
