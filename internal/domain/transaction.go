package domain

import (
	"time"

	"github.com/vladkaprelev/finance-go/internal/errs"
)

type Transaction struct {
	ID         uint      `gorm:"primaryKey;autoIncrement"`
	UserID     uint      `gorm:"index;not null"`
	Amount     float64   `gorm:"not null"`
	CategoryID uint      `gorm:"index;not null"`
	Date       time.Time `gorm:"not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func NewTransaction(
	categoryId uint,
	userId uint,
	amount float64,
	date time.Time,
) (*Transaction, error) {

	if userId == 0 {
		return nil, errs.NewValidationError("ID пользователя должен быть положительным числом")
	}

	if categoryId == 0 {
		return nil, errs.NewValidationError("ID категории должен быть положительным числом")
	}

	if amount <= 0 {
		return nil, errs.NewValidationError("сумма транзакции должна быть положительным числом")
	}

	if date.IsZero() {
		return nil, errs.NewValidationError("дата транзакции не может быть пустой")
	}

	return &Transaction{
		CategoryID: categoryId,
		UserID:     userId,
		Date:       date,
		Amount:     amount,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}, nil
}
