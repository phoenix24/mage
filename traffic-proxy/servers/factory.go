package servers

import (
	"traffic-proxy/common"
	"traffic-proxy/configs"
)

type Proxy interface {
	ListenAndServe() error
}

func NewProxy(config configs.ServerConfig, channel chan *common.Packet) (Proxy, error) {
	var proxy = &ProxyServer{
		sink:    channel,
		name:    config.Name,
		mode:    config.Mode,
		source:  config.Source,
		backend: config.Backend,
	}
	return proxy, nil
}
