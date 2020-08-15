package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"syscall"
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
		log.Println("forwarding the incoming request :", request)
	})
}

func NewServer(port string, backend string) {
	var mux = http.NewServeMux()
	mux.Handle("/", handler(backend))

	var server = &http.Server{Addr: port, Handler: mux}
	log.Fatalln(server.ListenAndServe())
}

func main() {
	var routes = os.Args[1]
	for index, route := range strings.Split(routes, ",") {
		var parts = strings.Split(route, ":")
		var src, dst = ":" + parts[0], "http://localhost:" + parts[1]

		log.Printf("%d, proxying: %s => %s\n", index, src, dst)
		go NewServer(src, dst)
	}

	var signals = make(chan os.Signal)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	for {
		<-signals
		log.Println("\r- Ctrl+C pressed in Terminal")
		os.Exit(0)
	}
}
