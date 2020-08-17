package protocols

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"traffic-proxy/configs"
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

func newService(backend configs.ServerConfig) {
	var mux = http.NewServeMux()
	mux.Handle("/", handler(backend.Backend))

	var server = &http.Server{Addr: port, Handler: mux}
	log.Fatalln(server.ListenAndServe())
}
