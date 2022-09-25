package main

import (
	"crypto/md5"
	"crypto/tls"
	"fmt"
	"net/http"
)

func main() {
	println("hello")
	s := &http.Server{
		Addr:    ":1234",
		Handler: nilHandler,
		TLSConfig: &tls.Config{
			GetConfigForClient: getClientConfig,
		},
	}
	s.ListenAndServeTLS("cert.pem", "key.pem")
}

var nilHandler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte("{}"))
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
