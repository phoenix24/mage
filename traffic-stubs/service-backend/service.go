package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"thetestr/traffic-stubs/service-commons"
)

func readConf(path string) SvcConf {
	viper.SetConfigType("yaml")

	var config SvcConf
	var content, _ = ioutil.ReadFile(path)
	var _ = viper.ReadConfig(bytes.NewBuffer(content))
	if  err := viper.Unmarshal(&config); err != nil {
		fmt.Println("failed to read service config")
	}
	return config
}

func handleInfo(writer http.ResponseWriter, request *http.Request) {
	var response, err = json.Marshal(commons.Info{"info", "v1"})
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(response)
}

func handleQuote(writer http.ResponseWriter, request *http.Request) {
	var response, err = json.Marshal(commons.Quote{
		"napoleon hill",
		"your big opportunity may be right where you are now.",
	})
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(response)
	log.Println("request served!")
}

func main() {
	var config = readConf(os.Args[1])
	http.HandleFunc("/info", handleInfo)
	http.HandleFunc("/quote", handleQuote)

	log.Println("service config :", config)
	log.Fatalln(http.ListenAndServe(config.HostPort(), nil))
}
