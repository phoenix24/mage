package servers

type MockServer struct {
}

func (m *MockServer) ListenAndServe() error {
	panic("implement me")
}
