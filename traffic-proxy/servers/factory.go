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
	var writers []io.Writer
	for _, sink := range config.Sinks {
		if sink == "null" {
			writers = append(writers, &sinks.NullSink{})
		}
		if sink == "console" {
			writers = append(writers, &sinks.ConsoleSink{})
		}
	}

	var server = &ProxyServer{
		name:    config.Name,
		mode:    config.Mode,
		sinks:   writers,
		source:  config.Source,
		backend: config.Backend,
	}
	return server, nil
}
