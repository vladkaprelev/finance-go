package repo

import (
	"time"

	"github.com/vladkaprelev/finance-go/internal/domain/model"
)

type IBudgetRepository interface {
	Create(budget *model.Budget) error
	Update(budget *model.Budget) error
	Delete(id uint) error

	FindByID(id uint) (*model.Budget, error)
	FindByUserID(userID uint, startDate time.Time, endDate time.Time) ([]*model.Budget, error)
	FindByCategoryID(categoryID uint, startDate time.Time, endDate time.Time) ([]*model.Budget, error)
}
