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
	var chPackets = make(chan *common.Packet, 100)
	var chCommands = make(chan *common.Commands, 100)

	var psink, _ = sinks.NewPacketSink(pconfig.Sinks, chPackets, chCommands)
	go psink.Consume()

	for _, config := range pconfig.Servers {
		go func(config configs.ServerConfig) {
			var server, _ = servers.NewProxy(config, chPackets, chCommands)
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
