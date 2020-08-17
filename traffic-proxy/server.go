package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"traffic-proxy/configs"
	"traffic-proxy/services"
	"traffic-proxy/storages"
)

func main() {
	var config = configs.ReadConfig(os.Args[1])
	var storage = storages.DataStore(config.Storage)
	for _, backend := range config.Services {
		var service = services.NewService(backend, storage)
		go service.ListenAndServe()
	}

	var signals = make(chan os.Signal)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	for {
		<-signals
		log.Println("\r- Ctrl+C pressed in Terminal")
		os.Exit(0)
	}
}
