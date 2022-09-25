package main

import (
	"crypto/md5"
	"crypto/tls"
	"fmt"
	"net/http"
)

func main() {
	port := 1234

	fmt.Printf("Server listens on port %d, it will print the client TLS fingerprint each time it receives a connection\n", port)

	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: nilHandler,
		TLSConfig: &tls.Config{
			GetConfigForClient: getClientConfig,
		},
	}
	err := s.ListenAndServeTLS("cert.pem", "key.pem")
	if err != nil {
		fmt.Printf("cannot start web server: %s\n", err)
		fmt.Println("Maybe you forgot to create the TLS certificates ?")
		fmt.Println("  openssl req -x509 -nodes -newkey rsa:4096 -keyout key.pem -out cert.pem -sha256 -days 365")
	}
}

var nilHandler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	_, _ = w.Write([]byte("{}"))
}

// called by go runtime when ClientHelloInfo is available
// next step is to attach the fingerprint to the connection to have the information available while processing the request
func getClientConfig(helloInfo *tls.ClientHelloInfo) (*tls.Config, error) {
	h := md5.New()
	h.Write([]byte(fmt.Sprintf("%v-%v-%v-%v-%v-%v\n",
		helloInfo.SupportedVersions,
		helloInfo.CipherSuites,
		helloInfo.SupportedCurves,
		helloInfo.SupportedPoints,
		helloInfo.SignatureSchemes,
		helloInfo.SupportedProtos,
	)))
	fingerprint := fmt.Sprintf("%x", h.Sum(nil))

	fmt.Printf("client config TLS fingerprint: %s\n", fingerprint)

	return nil, nil
}
