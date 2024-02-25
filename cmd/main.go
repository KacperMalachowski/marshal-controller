package main

import (
	"bufio"
	"changeme/pkg/td2"
	"context"
	"fmt"
	"log"
	"os"
)

func main() {
	addr := "127.0.0.1:7424"

	client := td2.New(context.Background(), "xd")

	client.Connect(addr)

	go func() {
		for {
			log.Print(<-client.ReadChan)
		}
	}()

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">")
		text, _ := reader.ReadString('\n')

		log.Println([]byte(text))
		client.Write(text)
	}
}
