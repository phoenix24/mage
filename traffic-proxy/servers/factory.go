package servers

import "traffic-proxy/configs"

type Server interface {
	ListenAndServe(config configs.ServerConfig) error
}

func NewServer(config configs.ServerConfig) (Server, error) {
	return &ProxyServer{}, nil
}
