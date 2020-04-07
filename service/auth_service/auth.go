package auth_service

import (
	"um_sys/module"
)

type Auth struct {
	Email    string
	Password string
}

func (a *Auth) Check() (bool, error) {
	return module.CheckAuth(a.Email, a.Password)
}

func (a *Auth) ExistEmail() (bool, error) {
	return module.ExistEmail(a.Email)

}

func (a *Auth) Register() (uint64, error) {
	return module.Register(a.Email, a.Password)
}
