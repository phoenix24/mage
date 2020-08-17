package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"traffic-proxy/configs"
	"traffic-proxy/servers"
)

func main() {
	var pconfig = configs.ReadConfig(os.Args[1])
	for _, config := range pconfig.Servers {
		var server, _ = servers.NewServer(config)
		go func(config configs.ServerConfig) {
			var _ = server.ListenAndServe(config)
		}(config)
	}

	var signals = make(chan os.Signal)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	for {
		<-signals
		log.Println("\r- Ctrl+C pressed in Terminal")
		os.Exit(0)
	}
}
