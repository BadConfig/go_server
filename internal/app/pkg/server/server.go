package server

type server struct{}

func New() *server {
	return &server{}
}

func (s *server) Start() error {
	return nil
}
