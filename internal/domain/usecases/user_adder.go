package usecases

import (
	"context"
)

//go:generate mockgen -destination=mocks/user_adder_mock.go -package=mocks . UserAdder
type UserAdder interface {
	AddUser(ctx context.Context, userID string) error
}
