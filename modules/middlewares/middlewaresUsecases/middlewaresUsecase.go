package middlewaresUsecases

import (
	"fmt"

	"github.com/muhammadfarhankt/nft-marketplace/modules/middlewares"
	"github.com/muhammadfarhankt/nft-marketplace/modules/middlewares/middlewaresRepositories"
)

type NMiddlewaresUsecase interface {
	FindAccessToken(userId, accessToken string) bool
	FindRole() ([]*middlewares.Role, error)
}

type middlewaresUsecase struct {
	middlewaresRepository middlewaresRepositories.NMiddlewaresRepository
}

func MiddlewaresUsecase(middlewaresRepository middlewaresRepositories.NMiddlewaresRepository) NMiddlewaresUsecase {
	return &middlewaresUsecase{
		middlewaresRepository: middlewaresRepository,
	}
}

func (m *middlewaresUsecase) FindAccessToken(userId, accessToken string) bool {
	return m.middlewaresRepository.FindAccessToken(userId, accessToken)
}

func (m *middlewaresUsecase) FindRole() ([]*middlewares.Role, error) {
	roles, err := m.middlewaresRepository.FindRole()
	fmt.Println(roles, err)
	if err != nil {
		return nil, err
	}
	return roles, nil
}
