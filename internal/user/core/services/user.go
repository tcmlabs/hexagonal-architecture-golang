package services

import (
	"fmt"

	"github.com/google/uuid"
	"tcmlabs.fr/hexagonal_architecture_golang/internal/user/core"
	"tcmlabs.fr/hexagonal_architecture_golang/internal/user/secondary_adapter/repositories/user"
)

type User interface {
	UserRetriever
	UserAppender
}

type userImpl struct {
	userRepository user.Repository
}

func NewUserServices(userRepository user.Repository) User {
	return &userImpl{
		userRepository: userRepository,
	}
}

type UserAppender interface {
	Create(email string) (*core.User, error)
}

type UserRetriever interface {
	Get() ([]core.User, error)
}

func (u userImpl) Get() ([]core.User, error) {
	userRetrieved, err := u.userRepository.Get()
	if err != nil {
		return nil, fmt.Errorf("user: failed to get user: %w", err)
	}

	return userRetrieved, nil
}

func (u userImpl) Create(email string) (*core.User, error) {
	var dUser = &core.User{
		Email: email,
		ID:    uuid.New().String(),
	}

	if err := u.userRepository.Create(dUser); err != nil {
		return nil, fmt.Errorf("user: failed to create user: %w", err)
	}

	return dUser, nil
}
