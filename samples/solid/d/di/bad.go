package main


type ServiceA struct{}

func NewServiceA() *ServiceA{
	return new(ServiceA)
}

type ServiceB struct{}

func NewServiceB() *ServiceB {
	return new(ServiceB)
}

type Server struct {
	serviceA *ServiceA
	serviceB *ServiceB
}

func NewServer() *Server{
	return &Server{NewServiceA(), NewServiceB()} // we create the dependencies inside the constructor
}

func (Server) start() {

}

func main() {
	server := NewServer()
	server.start()
}