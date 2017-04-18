package service

// Service is an interface for ...
type Service interface {
	Perform(req *Request) error
}

type service struct {
	Context
}

// New creates a new service instance from the context
func New(c Context) Service {
	return &service{
		Context: c,
	}
}

func (s *service) Perform(req *Request) error {
	return nil
}
