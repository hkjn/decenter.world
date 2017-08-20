package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/crypto/acme/autocert"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `<html><body><p>Hello! Not much here, so for now you should check out this YouTube video: <a href="https://www.youtube.com/watch?v=IrSn3zx2GbM">link</a></body>`)
	})
	addr := os.Getenv("DECENTER_WORLD_ADDR")
	if addr == "" {
		addr = ":8080"
	}
	s := &http.Server{
		Addr: addr,
	}
	if addr == ":443" {
		fmt.Println("Serving TLS..")
		m := autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			Cache:      autocert.DirCache("/etc/secrets/acme/"),
			HostPolicy: autocert.HostWhitelist("decenter.world"),
		}
		s.TLSConfig = &tls.Config{GetCertificate: m.GetCertificate}
		log.Fatal(s.ListenAndServeTLS("", ""))
	} else {
		fmt.Printf("Serving plaintext HTTP on %s..\n", addr)
		log.Fatal(s.ListenAndServe())
	}
}
