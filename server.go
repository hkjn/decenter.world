package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `<html><body><p>Hello! Not much here, so for now you should check out this YouTube video: <a href="https://www.youtube.com/watch?v=IrSn3zx2GbM">link</a></body>`)
	})
	addr := os.Getenv("DECENTER_WORLD_ADDR")
	if addr == "" {
		addr = ":8080"
	}
	log.Fatal(http.ListenAndServe(addr, nil))
}
