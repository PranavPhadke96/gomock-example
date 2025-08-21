package greeting

type Greeter interface {
	Greet(name string) string
}

type GreetingService struct {
	greeter Greeter
}

func NewGreetingService(g Greeter) *GreetingService {
	return &GreetingService{greeter: g}
}

func (s *GreetingService) CreateGreeting(name string) string {
	return s.greeter.Greet(name)
}