pcakage main 

import "net"


type Confir struct {
	ListenAddress string
}

type server struct {
	Config
	ln net.Listener
}

func main() {

}