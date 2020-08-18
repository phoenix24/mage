package servers

import "traffic-proxy/configs"

type Server interface {
	ListenAndServe() error
}

func NewServer(config configs.ServerConfig) (Server, error) {
	var server = &ProxyServer{
		name:    config.Name,
		mode:    config.Mode,
		source:  config.Source,
		backend: config.Backend,
		storage: config.Storage,
	}
	return server, nil
}
