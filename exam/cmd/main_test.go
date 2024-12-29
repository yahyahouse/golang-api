// Server starts successfully and listens for incoming requests
package main_test

import (
	"exam/internal/provider"
	"net"
	"net/http"
	"testing"
	"time"
)

func TestServerStartsSuccessfully(t *testing.T) {
	prov := provider.GetProvider()
	server := prov.GetServer()

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			t.Fatalf("Expected server to start successfully, got error: %v", err)
		}
	}()

	time.Sleep(1 * time.Second) // Allow some time for the server to start
	if server.Addr == "" {
		t.Fatal("Expected server to have a valid address")
	}
}
func TestServerFailsToStartDueToPortBindingIssues(t *testing.T) {
	prov := provider.GetProvider()
	server := prov.GetServer()

	// Simulate port binding issue by starting another server on the same port
	conflictingServer := &http.Server{Addr: server.Addr}
	go conflictingServer.ListenAndServe()
	defer conflictingServer.Close()

	errChan := make(chan error)
	go func() {
		if err := server.ListenAndServe(); err != nil {
			errChan <- err
		}
	}()

	select {
	case err := <-errChan:
		if err == nil {
			t.Fatal("Expected server to fail due to port binding issues, but it started successfully")
		}
	}
}
func TestServerFailsToStartPortInUse(t *testing.T) {
	prov := provider.GetProvider()
	server := prov.GetServer()

	// Simulate port in use by listening on the same port
	listener, err := net.Listen("tcp", server.Addr)
	if err != nil {
		t.Fatalf("Failed to listen on port: %v", err)
	}
	defer listener.Close()

	errChan := make(chan error)
	go func() {
		if err := server.ListenAndServe(); err != nil {
			errChan <- err
		}
	}()

	select {
	case err := <-errChan:
		if err == nil {
			t.Fatal("Expected error due to port being in use, got nil")
		}
	}
}
