package security

type baseMapper interface {
	AddUser(login string, password string) error
}

// Registrator интерфейс сервиса регистрации
type Registrator interface {
	AddUser(login string, password string) (err error)
}

type registrator struct {
	userBaseMapper baseMapper
}

func (r *registrator) AddUser(login string, password string) (err error) {
	err = r.userBaseMapper.AddUser(login, password)
	return
}

// NewRegistrator конструктор сервиса регистрации
func NewRegistrator(userBaseMapper baseMapper) Registrator {
	return &registrator{
		userBaseMapper: userBaseMapper,
	}
}
