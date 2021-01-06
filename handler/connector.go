package handler

import (
	"github.com/FerdinaKusumah/fastcrud/repository"
	"github.com/FerdinaKusumah/fastcrud/usecase"
	"sync"
)

var (
	useCaseOnce sync.Once
	uc          = new(usecase.UseCase)
)

func GetUseCase() *usecase.UseCase {
	useCaseOnce.Do(func() {
		uc = usecase.NewUseCase(
			getRepository(),
		)
	})
	return uc
}

var (
	repo    *repository.Repository
	oneRepo sync.Once
)

func getRepository() *repository.Repository {
	oneRepo.Do(func() {
		repo = repository.NewRepository()
	})

	return repo
}
