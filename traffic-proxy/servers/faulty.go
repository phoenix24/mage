package servers

type FaultyServer struct {
}

func (f *FaultyServer) ListenAndServe() error {
	panic("implement me")
}
