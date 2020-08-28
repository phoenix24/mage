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
		sink:   channel,
		name:   config.Name,
		mode:   config.Mode,
		source: config.Source,
		remote: config.Remote,
		protocol: config.Remote.Protocol(),
	}
	return proxy, nil
}
