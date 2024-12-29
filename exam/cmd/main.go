package main

import (
	"context"
	"exam/internal/provider"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var prov = provider.GetProvider()

func main() {

	server := prov.GetServer()
	errChan := make(chan error)
	go func() {
		log.Println("[Main] server.ListenAndServe()")
		if err := server.ListenAndServe(); err != nil {
			errChan <- err
		}
	}()

	waitOSSignal(server, errChan)
}
func waitOSSignal(server *http.Server, errChan <-chan error) {
	//Setting up Interrupt signal capturing
	term := make(chan os.Signal, 1)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)
	log.Println("[Main] waiting server done or signal")
	select {
	case <-term: // Waiting for SIGINT (pkill -2) or SIGTERM
		log.Println("[Main] Signal terminate detected")
		shutdownServer(server)
	case err := <-errChan: // error listening
		log.Fatalf("[Main] Server error: %v", err)
		panic(err)
	}
}

func shutdownServer(server *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("[Main] Error Shutdown Server %v", err)
		panic(err)
	}
}
