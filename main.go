package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const port = 8080

func main() {
	ctx := context.Background()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	srv := &http.Server{Addr: fmt.Sprintf(":%d", port)}
	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
			fmt.Fprintf(w, "Hello, World!")
		})

		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("failed to start server: %v", err)
		}
	}()

	<-signals
	log.Printf("shutting down server")
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("failed to shutdown server: %v", err)
	}
}
