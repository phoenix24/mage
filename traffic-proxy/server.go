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

	//initialize all psinks.
	var psinks []*sinks.TrafficSink
	for _, config := range pconfig.Sinks {
		var sink, _ = sinks.NewSink(config, channel)
		psinks = append(psinks, sink)
	}

	//initialize sink fanout.
	sinks.NewSinkFanout(channel, psinks)

	//initialize all proxies.
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
