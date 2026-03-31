package server

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/acme/autocert"
)

type Server struct {
	Handler http.Handler
	Domain  string
}

func NewServer(handler http.Handler, domain string) *Server {
	return &Server{
		Handler: handler,
		Domain:  domain,
	}
}

func (s *Server) Start() {
	// Load .env
	_ = godotenv.Load()

	localMode := os.Getenv("LOCAL") == "true"

	if localMode {
		// LOCAL HTTP mode
		log.Println("Starting server in LOCAL HTTP mode on :8080")
		log.Fatal(http.ListenAndServe(":8080", s.Handler))
		return
	}

	// PRODUCTION HTTPS mode
	certManager := &autocert.Manager{
		Cache:      autocert.DirCache("certs"),
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(s.Domain),
	}

	go func() {
		http.ListenAndServe(":80", certManager.HTTPHandler(nil))
	}()

	server := &http.Server{
		Addr:      ":443",
		Handler:   s.Handler,
		TLSConfig: certManager.TLSConfig(),
	}

	log.Printf("HTTPS server running on :443 for domain %s\n", s.Domain)
	log.Fatal(server.ListenAndServeTLS("", ""))
}
