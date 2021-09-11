package auth

import (
	"errors"

	"github.com/rschio/grpc-poc/user"
)

type Auther interface {
	Auth(token, method string) (*user.User, error)
}

type MyAuther struct{}

func (a MyAuther) Auth(token, method string) (*user.User, error) {
	publicMethods := []string{"/proto.Tracker/Login"}
	for _, m := range publicMethods {
		if method == m {
			return nil, nil
		}
	}

	const validToken = "Bearer consegue"
	if token != validToken {
		return nil, errors.New("invalid token")
	}
	return &user.User{Name: "Moises"}, nil
}
