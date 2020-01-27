package authentication

import "errors"

type baseMapper interface {
	GetUserPassword(login string, password string) (basePassword string, err error)
}

type registrator interface {
	AddUser(login string, password string) error
}

type service interface {
	Connect(login string, msg *[]byte) error
}

// Authentificator интерфейс аутентификации
type Authentificator interface {
	ConnectUser(login string, password string, msg *[]byte) (err error)
}

type authentificator struct {
	userBaseMapper  baseMapper
	userRegistrator registrator
	userService     service
}

func (a *authentificator) ConnectUser(login string, password string, msg *[]byte) (err error) {
	savedPass, baseErr := a.userBaseMapper.GetUserPassword(login, password)
	if baseErr != nil {
		err = a.userRegistrator.AddUser(login, password)
		if err != nil {
			return
		}
		savedPass = password
	}
	if savedPass != password {
		err = errors.New("wrong password")
		return
	}
	err = a.userService.Connect(login, msg)
	return
}

// NewAuthentificator конструктор аутентефикатора пользователей
func NewAuthentificator(userBaseMapper baseMapper, userRegistrator registrator, userService service) Authentificator {
	return &authentificator{
		userBaseMapper:  userBaseMapper,
		userRegistrator: userRegistrator,
		userService:     userService,
	}
}
