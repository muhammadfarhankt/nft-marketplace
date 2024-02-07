package middlewaresUsecases

import "github.com/muhammadfarhankt/nft-marketplace/modules/middlewares/middlewaresRepositories"

type NMiddlewaresUsecase interface {
}

type middlewaresUsecase struct {
	middlewaresRepository middlewaresRepositories.NMiddlewaresRepository
}

func MiddlewaresUsecase(middlewaresRepository middlewaresRepositories.NMiddlewaresRepository) NMiddlewaresUsecase {
	return &middlewaresUsecase{
		middlewaresRepository: middlewaresRepository,
	}
}
