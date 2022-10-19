package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	log.Println(http.ListenAndServe(":6060", nil))
}

// go tool pprof -http :5000 localhost:6060/debug/pprof/profile
