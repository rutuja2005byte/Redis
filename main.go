pcakage main 

import "net"


type Confir struct {
	ListenAddress string
}

type server struct {
	Config
	ln net.Listener
}

func NewServer(cfg config) *server {
	return &Server{
		config: cfg,
	}
}

func (s *Server) start() error {
	ln,
}

func main() {

}