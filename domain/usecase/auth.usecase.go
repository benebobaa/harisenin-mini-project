package usecase

import (
	"github.com/benebobaa/harisenin-mini-project/domain/entity"
	"github.com/benebobaa/harisenin-mini-project/domain/repository"
	"github.com/benebobaa/harisenin-mini-project/shared/util"
)

type AuthUsecase interface {
	Login(action entity.User) (*entity.User, error)
	Register(action entity.User) error
}

type authUsecaseImpl struct {
	authRepository repository.AuthRepository
}

func NewAuthUsecase(authRepository repository.AuthRepository) authUsecaseImpl {
	return authUsecaseImpl{authRepository: authRepository}
}

func (a *authUsecaseImpl) Login(action entity.User) (*entity.User, error) {
	var user entity.User

	err := a.authRepository.FindByUsername(&user, action.Username)
	if err != nil {
		return nil, err
	}

	err = util.CheckPassword(action.Password, user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (a *authUsecaseImpl) Register(action entity.User) error {

	hashedPassword, err := util.HashPassword(action.Password)
	if err != nil {
		return err
	}

	action.Password = hashedPassword

	err = a.authRepository.CreateUser(&action)

	if err != nil {
		return err
	}

	return nil
}
