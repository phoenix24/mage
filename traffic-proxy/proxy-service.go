package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func handler(backend string) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		var pUrl, _ = url.Parse(backend)
		var revproxy = httputil.NewSingleHostReverseProxy(pUrl)
		request.Host = pUrl.Host
		request.URL.Host = pUrl.Host
		request.URL.Scheme = pUrl.Scheme
		request.Header.Set("X-Forwarded-Host", pUrl.Host)
		revproxy.ServeHTTP(writer, request)
	})
}

func main() {
	var backend = os.Args[1]
	var signals = make(chan os.Signal)

	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-signals
		log.Println("\r- Ctrl+C pressed in Terminal")
		os.Exit(0)
	}()

	var mux = http.NewServeMux()
	mux.Handle("/", handler(backend))

	var server = &http.Server{
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 20 * time.Second,
		IdleTimeout:  120 * time.Second,
		Addr:         ":7070",
		Handler:      mux,
	}
	log.Fatal(server.ListenAndServe())
}
