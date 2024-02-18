package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	addr := "127.0.0.1:7424"

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	go read(conn)

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')

		if text == "quit" {
			os.Exit(0)
		}

		conn.SetWriteDeadline(time.Now().Add(1 * time.Second))
		_, err := conn.Write([]byte(text))
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}

func read(conn net.Conn) {
	for {
		scanner := bufio.NewScanner(conn)

		for {
			ok := scanner.Scan()
			text := scanner.Text()

			fmt.Println(text)
			if !ok {
				fmt.Println("Reached EOF on server connection.")
				break
			}
		}
	}
}
