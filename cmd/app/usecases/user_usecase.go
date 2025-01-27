package usecases

import (
	"github.com/vladkaprelev/finance-go/internal/domain/model"
	"github.com/vladkaprelev/finance-go/internal/domain/repo"
	"github.com/vladkaprelev/finance-go/internal/errs"
)

type IUserUsecase interface {
	RegisterUser(email, password string) (*model.User, error)
	GetUserByID(id uint) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
}

type userUsecase struct {
	userRepo repo.IUserRepository
}

func NewUserUseCase(repo repo.IUserRepository) IUserUsecase {
	return &userUsecase{
		userRepo: repo,
	}
}

func (u *userUsecase) RegisterUser(email, password string) (*model.User, error) {
	if email == "" || password == "" {
		return nil, errs.NewValidationError("RegisterUser: Некорректные данные пользователя")
	}

	user := &model.User{
		Email:    email,
		Password: password,
	}

	createdUser, err := u.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (u *userUsecase) GetUserByID(id uint) (*model.User, error) {
	user, err := u.userRepo.FindByID(id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userUsecase) GetUserByEmail(email string) (*model.User, error) {
	user, err := u.userRepo.FindByEmail(email)

	if err != nil {
		return nil, err
	}

	return user, nil
}
