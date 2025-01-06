package domain

import (
	"time"

	"github.com/vladkaprelev/finance-go/internal/errs"
)

type Transaction struct {
	ID         uint `gorm:"primaryKey;autoIncrement"`
	UserID     uint
	CategotyID uint
	Value      int32
	Date       time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func NewTransaction(
	categoryId uint,
	userId uint,
	value int32,
	date time.Time,
) (*Transaction, error) {

	if userId == 0 {
		return nil, errs.NewValidationError("ID пользователя должен быть положительным числом")
	}

	if categoryId == 0 {
		return nil, errs.NewValidationError("ID категории должен быть положительным числом")
	}

	if value < 0 {
		return nil, errs.NewValidationError("значение value не может быть отрицательным")
	}

	return &Transaction{
		CategotyID: categoryId,
		UserID:     userId,
		Date:       date,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}, nil
}
