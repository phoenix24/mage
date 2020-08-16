package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"thetestr/traffic-stubs/service-commons"
)

func main () {
	var port = ":" + os.Args[1]
	var backends = strings.Split(os.Args[2], ",")
	log.Println("listening on port", port)
	log.Println("listening on backends", backends)

	http.Handle("/infos", handleInfos(backends))
	http.Handle("/quotes", handleQuotes(backends))

	log.Println("frontend listening on", port)
	log.Fatalln(http.ListenAndServe(port, nil))
}


func fetch(index int, backend string, v interface{}) ([]byte, error) {
	log.Printf("fetching from %d, %s\n", index, backend)
	var response, _ = http.Get(backend)
	defer response.Body.Close()

	var payload, _ = ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(payload, &v)
	return payload, nil
}

func handleInfos(backends []string) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, rq *http.Request) {
		var infos []commons.Info
		for index, backend := range backends {
			var info commons.Info
			var _, _ = fetch(index, "http://"+backend+"/info", &info)
			infos = append(infos, info)
		}

		var response, _ = json.Marshal(commons.Infos{len(infos), infos})
		wr.Header().Set("Content-Type", "application/json")
		wr.Write(response)
	})
}

func handleQuotes(backends []string) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, rq *http.Request) {
		var quotes []commons.Quote
		for index, backend := range backends {
			var quote commons.Quote
			var _, _ = fetch(index, "http://"+backend+"/quote", &quote)
			quotes = append(quotes, quote)
		}

		var response, _ = json.Marshal(commons.Quotes{len(quotes), quotes})
		wr.Header().Set("Content-Type", "application/json")
		wr.Write(response)
	})
}

