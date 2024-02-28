package td2

import (
	"context"
	"net"
	"testing"
)

func TestConnectionSuccessfull(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Errorf("cannot setup server: %s", err)
	}

	go func() {
		defer ln.Close()
		_, err = ln.Accept()
		if err != nil {
			t.Errorf("failed to accept connection: %s", err)
		}
	}()

	td2Client := New(context.Background(), "someHash")
	err = td2Client.Connect(ln.Addr().String())
	if err != nil {
		t.Error(err)
	}
}

func TestConnectionFailed(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Errorf("cannot setup server: %s", err)
	}

	go func() {
		defer ln.Close()
	}()

	td2Client := New(context.Background(), "someHash")
	err = td2Client.Connect(ln.Addr().String())
	if err == nil {
		t.Error("got no error when expected to")
	}
}
