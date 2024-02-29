package td2_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/kacpermalachowski/marshal-controller/internal/testserver"
	"github.com/kacpermalachowski/marshal-controller/pkg/td2"
)

func TestNewClient(t *testing.T) {
	ctx := context.Background()
	stationHash := "exampleHash"
	client := td2.New(ctx, stationHash)

	if client.IsConnected {
		t.Error("Expected IsConnected to be false, but got true")
	}

	if len(client.ReadChan) != 0 {
		t.Error("Expected ReadChan to be empty, but it is not")
	}
}

func TestConnectDisconnect(t *testing.T) {
	ctx := context.Background()
	stationHash := "exampleHash"
	client := td2.New(ctx, stationHash)

	server := testserver.New()
	addr, err := server.Start()
	if err != nil {
		t.Fatalf("Failed to setup test server: %s", err)
	}

	// Test Connect
	err = client.Connect(addr)
	if err != nil {
		t.Errorf("Expected no error on Connect, but got: %v", err)
	}

	if !client.IsConnected {
		t.Error("Expected IsConnected to be true after Connect, but got false")
	}

	// Test Disconnect
	client.Disconnect()
	if client.IsConnected {
		t.Error("Expected IsConnected to be false after Disconnect, but got true")
	}
}

func TestWrite(t *testing.T) {
	ctx := context.Background()
	stationHash := "exampleHash"
	client := td2.New(ctx, stationHash)
	testData := "Hello, World!"
	expectedData := fmt.Sprintf("%s\r\n", testData)

	server := testserver.New()
	addr, err := server.Start()
	if err != nil {
		t.Fatalf("Failed to setup test server: %s", err)
	}

	err = client.Connect(addr)
	if err != nil {
		t.Errorf("Failed to connect: %s", err)
	}
	defer client.Disconnect()

	err = client.Write(testData)
	if err != nil {
		t.Errorf("Expected no error on Write operation, but got: %s", err)
	}

	receivedData := <-server.ReceivedDataChan
	if receivedData != expectedData {
		t.Errorf("Expected written data to be %q, but got %q", expectedData, receivedData)
	}
}
