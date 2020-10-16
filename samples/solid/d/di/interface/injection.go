package main

type ServiceA struct{}

func NewServiceA() *ServiceA {
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

// Server has a published interface
type ServiceSetterInterface interface {
	SetServiceA(sa *ServiceA)
	SetServiceB(sb *ServiceB)
	start()
}

func NewServer() *Server {
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
	var server ServiceSetterInterface // Notice the subtle difference here: the injector does not really need to now the concrete type. it only needs the interface.
	server = NewServer()
	// Dependencies are constructed
	serviceA := NewServiceA()
	serviceB := NewServiceB()
	// Dependencies are injected in a later step, after object creation
	server.SetServiceA(serviceA)
	server.SetServiceB(serviceB)

	server.start()
}
