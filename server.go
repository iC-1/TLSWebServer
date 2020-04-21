package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"time"
)

// RunTLSServer configures and starts a TLS web server
func (app *App) RunTLSServer() {
	// define the TLS config
	tlsConfig := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		MaxVersion:               tls.VersionTLS12,
		PreferServerCipherSuites: true,
		CurvePreferences:         []tls.CurveID{tls.X25519, tls.CurveP256},
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,
			tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
		},
	}

	// Add Idle, Read and Write timeouts to the server.
	srv := &http.Server{
		Addr:         app.domain+":"+app.httpsport,
		Handler:      app.Routes(),
		TLSConfig:    tlsConfig,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	// make sure we can reload certificates and keys during runtime
	kpr, err := NewKeypairReloader(cfg.TLSCertPath, cfg.TLSKeyPath)
	if err != nil {
		log.Fatal(err)
	}
	srv.TLSConfig.GetCertificate = kpr.GetCertificateFunc()
	// start the server
	//log.Printf("Starting server on %s", app.httpsport)
	err = srv.ListenAndServeTLS("", "")
	log.Fatal(err)
}
