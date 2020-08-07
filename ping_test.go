package hello

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingOrder(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	}))
	defer server.Close()

	config := &Config{HttpClient: server.Client()}
	client, err := NewClient(config)
	if err != nil {
		t.Fatal(err)
	}
	client.PingOrder()
}
