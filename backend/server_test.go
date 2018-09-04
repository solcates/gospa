package backend

import (
	"net/http"
	"testing"
)

func TestRunServer(t *testing.T) {
	httpListenAndServe = func(addr string, handler http.Handler) error {
		return nil
	}
	tests := []struct {
		name string
	}{
		{
			name: "Default",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RunServer()
		})
	}
}
