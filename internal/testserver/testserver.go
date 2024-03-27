package testserver

import (
	"fmt"
	"net"
)

type TestTCPServer struct {
	ReceivedDataChan chan string
}

func New() *TestTCPServer {
	return &TestTCPServer{
		ReceivedDataChan: make(chan string),
	}
}

func (s *TestTCPServer) Start() (string, error) {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return "", fmt.Errorf("failed to start test server: %s", err)
	}

	go func() {
		defer l.Close()

		for {
			conn, err := l.Accept()
			if err != nil {
				return
			}

			go func(c net.Conn) {
				defer c.Close()

				buf := make([]byte, 1024)
				n, err := c.Read(buf)
				if err != nil {
					return
				}

				s.ReceivedDataChan <- string(buf[:n])
			}(conn)
		}
	}()

	return l.Addr().String(), nil
}

func (s *TestTCPServer) Stop() {
	close(s.ReceivedDataChan)
}
