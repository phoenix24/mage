package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"traffic-proxy/common"
	"traffic-proxy/configs"
	"traffic-proxy/servers"
	"traffic-proxy/sinks"
)

func main() {
	var pconfig = configs.ReadConfig(os.Args[1])
	var channel = make(chan *common.Packet, 100)

	var psink, _ = sinks.NewPacketSink(pconfig.Sinks, channel)
	go psink.Consume()

	for _, config := range pconfig.Servers {
		go func(config configs.ServerConfig) {
			var server, _ = servers.NewProxy(config, channel)
			var _ = server.ListenAndServe()
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
