package customer

import "errors"

// Mapper интерфейс работы с базой данных
type Mapper interface {
	GetUserPassword(login string, password string) (basePassword string, err error)
	AddUser(login string, password string) (err error)
}

type userBase struct {
	userLoginPass map[string]string
}

func (u *userBase) GetUserPassword(login string, password string) (basePassword string, err error) {
	if savedPass, ok := u.userLoginPass[login]; ok && (savedPass == password) {
		return savedPass, nil
	}
	err = errors.New("pair login-password not found")
	return
}

func (u *userBase) AddUser(login string, password string) (err error) {
	if _, ok := u.userLoginPass[login]; ok {
		err = errors.New("user with tis logn already exist")
		return
	}
	u.userLoginPass[login] = password
	return
}

// NewUserBase возвращает интерфейс для взаимодействия с базой
func NewUserBase() Mapper {
	return &userBase{
		userLoginPass: make(map[string]string),
	}
}
