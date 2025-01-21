package repo

import (
	"time"

	"github.com/vladkaprelev/finance-go/internal/domain/model"
)

type ITransactionRepository interface {
	Create(budget *model.Transaction) error
	Update(budget *model.Transaction) error
	Delete(id uint) error

	FindByID(id uint) (*model.Transaction, error)
	FindByUserID(userID uint, startDate time.Time, endDate time.Time) ([]*model.Transaction, error)
	FindByCategoryID(categoryID uint, startDate time.Time, endDate time.Time) ([]*model.Transaction, error)
}
