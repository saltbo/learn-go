package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
)

func main() {
	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method, r.URL)
		if r.Method == http.MethodConnect {
			handleConnect(w, r)
			return
		}

		handlePlainHTTP(w, r)
	}))
}

func handlePlainHTTP(w http.ResponseWriter, r *http.Request) {
	req := r.Clone(context.Background())
	req.RequestURI = ""
	// TODO: REMOVE SOME HEADERS

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, err.Error(), resp.StatusCode)
		return
	}

	w.WriteHeader(resp.StatusCode)
	for k, vs := range resp.Header {
		for _, v := range vs {
			w.Header().Add(k, v)
		}
	}
	io.Copy(w, resp.Body)
}

func handleConnect(w http.ResponseWriter, r *http.Request) {
	destConn, err := net.Dial("tcp", r.URL.Host)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	w.WriteHeader(http.StatusOK)
	hijacker, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "Hijacking not supported", http.StatusInternalServerError)
		return
	}
	clientConn, _, err := hijacker.Hijack()
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	go func() {
		defer clientConn.Close()
		defer destConn.Close()
		_, err = io.Copy(clientConn, destConn)
		if err != nil {
			log.Printf("Error copying response to client: %s", err.Error())
		}
	}()
	go func() {
		_, err = io.Copy(destConn, clientConn)
		if err != nil {
			log.Printf("Error copying request to target: %s", err.Error())
		}
	}()
}
