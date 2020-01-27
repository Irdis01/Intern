package service

// Service интерфейс, описывающий работу сервиса
type Service interface {
	Connect(login string, msg *[]byte) (err error)
}

type service struct{}

func (s *service) Connect(login string, msg *[]byte) (err error) {
	serviceResponse := "Hi " + login
	newMsg := []byte(serviceResponse)
	*msg = append(*msg, newMsg[0:]...)
	return
}

// NewService конструктор новго сервиса
func NewService() Service {
	return &service{}
}
