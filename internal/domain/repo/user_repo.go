package repo

import "github.com/vladkaprelev/finance-go/internal/domain/model"

type IUserRepository interface {
	Create(user *model.User) error
	FindByID(id uint) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	Delete(id uint) error
}
