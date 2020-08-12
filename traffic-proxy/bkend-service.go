package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func handleInfo(writer http.ResponseWriter, request *http.Request) {
	var response, err = json.Marshal(Info{"info", "v1"})
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(response)
}

func handleQuote(writer http.ResponseWriter, request *http.Request) {
	var response, err = json.Marshal(Quote{
		"napoleon hill",
		"your big opportunity may be right where you are now.",
	})
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(response)
}

func main() {
	var port = ":" + os.Args[1]
	http.HandleFunc("/info", handleInfo)
	http.HandleFunc("/quote", handleQuote)

	log.Println("backend listening on", port)
	log.Fatalln(http.ListenAndServe(port, nil))
}
