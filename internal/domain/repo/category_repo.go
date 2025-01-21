package repo

import "github.com/vladkaprelev/finance-go/internal/domain/model"

type ICategoryRepository interface {
	Create(category *model.Category) error
	FindByID(id uint) (*model.Category, error)
	Update(category *model.Category) error
	Delete(id uint) error

	FindByUserID(userID uint) ([]*model.Category, error)
}
