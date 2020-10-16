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

func NewServer(a *ServiceA, b *ServiceB) *Server{
	return &Server{a, b} // Dependencies are injected in the constructor
}

func (Server) start() {

}

func main() {
	// Dependencies are constructed
	serviceA := NewServiceA()
	serviceB := NewServiceB()
	// Then they are injected, function main is the injector in this case
	server := NewServer(serviceA, serviceB)
	server.start()
}