package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/http2"
)

func loadClientCAs() *x509.CertPool {
	clientCAs := x509.NewCertPool()
	caCert, err := os.ReadFile("cert.pem")
	if err != nil {
		log.Fatalln("Could not load clientCA", err)
	}

	clientCAs.AppendCertsFromPEM(caCert)

	return clientCAs
}

func main() {

	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		logRequestDetails(r)
		fmt.Fprintf(w, "Handling incoming orders")
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		logRequestDetails(r)
		fmt.Fprintf(w, "Handling users")
	})

	port := 3000

	// Load the TLS cert and key
	// ---> to generate the certificate and key, we use the command ---> openssl req -x509 -newkey rsa:2048 -nodes -keyout key.pem -out cert.pem -days 365
	
	cert := "cert.pem"
	key := "key.pem"

	// Configure TLS
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
		ClientAuth: tls.RequireAndVerifyClientCert, // mTLS (mutual TLS)

		// in mTLS both client and server provide certificates to verify each other, thus this adds an extra layer of security
		// server verifies the certificate so only authorized clients can connect
		// say if we create an app, that doesnt require a browser, in that case we can use own self signed certificate but they also need to be verified

		ClientCAs: loadClientCAs(),
	}

	// Create a custom server
	server := &http.Server{
		Addr:      fmt.Sprintf(":%d", port),
		Handler:   nil,
		TLSConfig: tlsConfig,
	}

	// enable http2
	http2.ConfigureServer(server, &http2.Server{})

	fmt.Println("Server is running on port:", port)

	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Error starting server:", err)
	}

}

func logRequestDetails(r *http.Request) {
	httpVersion := r.Proto
	fmt.Println("Recieved request with http version:", httpVersion)

	if r.TLS != nil {
		tlsVersion := getTLSVersion(r.TLS.Version)
		fmt.Println("Recieved request with TLS version:", tlsVersion)
	} else {
		fmt.Println("Recieved request without TLS")
	}
}

func getTLSVersion(version uint16) string {
	switch version {
	case tls.VersionTLS10:
		return "TLS 1.0"
	case tls.VersionTLS11:
		return "TLS 1.1"
	case tls.VersionTLS12:
		return "TLS 1.2"
	case tls.VersionTLS13:
		return "TLS 1.3"
	default:
		return "Unknown TLS version"
	}
}

// Alternate way to generate key and certificate is to generate them separetely
// ---> we first need to generate the key, we use the command ---> openssl genpkey -algorithm RSA -out server.key -pkeyopt rsa_keygen_bits:2048
// ---> now we generate the certificate, we use the command ---> openssl req -new -x509 -key server.key -out server.crt -days 365