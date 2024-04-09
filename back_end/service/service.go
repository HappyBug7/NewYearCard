package service

type Service struct {
	Card
	DES
}

func New() *Service {
	service := &Service{}
	return service
}
