package servers

import (
	"traffic-proxy/configs"
)

type FaultyServer struct {
}

func (f *FaultyServer) ListenAndServe(config configs.ServerConfig) error {
	panic("implement me")
}
