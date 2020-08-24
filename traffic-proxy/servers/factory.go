package servers

import (
	"io"
	"traffic-proxy/configs"
	"traffic-proxy/sinks"
)

type Server interface {
	ListenAndServe() error
}

func NewServer(config configs.ServerConfig) (Server, error) {
	var tfsinks []io.Writer
	for _, sconfig := range config.Sinks {
		var tfsink, _ = sinks.NewSink(sconfig)
		tfsinks = append(tfsinks, tfsink)
	}

	var server = &ProxyServer{
		name:    config.Name,
		mode:    config.Mode,
		sinks:   tfsinks,
		source:  config.Source,
		backend: config.Backend,
	}
	return server, nil
}
