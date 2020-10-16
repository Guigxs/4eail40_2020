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
	return &Server{}
}

func (s *Server) SetServiceA(sa *ServiceA) {
	s.serviceA = sa
}

func (s *Server) SetServiceB(sb *ServiceB) {
	s.serviceB = sb
}

func (Server) start() {

}

func main() {
	// Server is created
	server := NewServer()
	// Dependencies are constructed
	serviceA := NewServiceA()
	serviceB := NewServiceB()
	// Dependencies are injected in a later step, after object creation
	server.SetServiceA(serviceA)
	server.SetServiceB(serviceB)

	server.start()
}