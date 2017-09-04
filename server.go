package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/crypto/acme/autocert"
)

const Page = `<html>
<body>
  <h1>decenter</h1>
  <i>Transitive verb</i>
  <blockquote>
  to cause to lose or shift from an established center or focus
  </blockquote>

  <iframe width="1120" height="630" src="https://www.youtube-nocookie.com/embed/IrSn3zx2GbM?showinfo=0&amp;start=60" frameborder="0" allowfullscreen></iframe>
</body>
</html>
`

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, Page)
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
