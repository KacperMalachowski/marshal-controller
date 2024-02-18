package td2

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type Client struct {
	conn        net.Conn
	stationHash []byte
	ReadChan    chan string
}

func New(stationHash string) *Client {
	return &Client{
		stationHash: []byte(stationHash),
		ReadChan:    make(chan string),
	}
}

func (c *Client) Connect(address string) error {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return err
	}

	go c.readMessages()

	c.conn = conn
	return nil
}

func (c *Client) Disconnect() {
	c.conn.Close()
}

func (c *Client) Write(data []byte) error {
	_, err := c.conn.Write(data)
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
		c.Write([]byte(fmt.Sprintf("/SHA2 %s", c.stationHash)))
	}
}
