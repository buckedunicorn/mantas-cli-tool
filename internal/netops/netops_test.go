package netops

import (
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingServer(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("Failed to listen on a port: %s", err)
	}
	defer listener.Close()

	go func() {
		conn, err := listener.Accept()
		if err != nil {
			t.Logf("Accept error: %s", err)
			return
		}
		conn.Close()
	}()

	PingServer(listener.Addr().String())
}

func TestFetchURL(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	FetchURL(server.URL)
}
