package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
)

type Handler interface {
	ServeHTTP(ctx context.Context, w http.ResponseWriter, r *http.Request)
}

type HandlerFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request)

func (h HandlerFunc) ServeHTTP(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	h(ctx, w, r)
}

func main() {
	abc := HandlerFunc(func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		http.DefaultServeMux.ServeHTTP(w, r)
	})

	http.HandleFunc("/aaa", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("abc"))
	})

	srv := &http.Server{
		Addr: ":9001",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			abc.ServeHTTP(context.Background(), w, r)
		}),
	}

	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		// We received an interrupt signal, shut down.
		if err := srv.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}

	<-idleConnsClosed
}
