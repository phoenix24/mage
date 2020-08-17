package servers

import (
	"traffic-proxy/configs"
)

type MockServer struct {
}

func (m *MockServer) ListenAndServe(config configs.ServerConfig) error {
	panic("implement me")
}
