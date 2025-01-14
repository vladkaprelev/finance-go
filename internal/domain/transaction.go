package domain

import (
	"time"

	"github.com/vladkaprelev/finance-go/internal/errs"
)

// Transaction представляет транзакцию пользователя, включающую информацию о сумме, категории,
// дате транзакции, а также временные метки создания и обновления.
type Transaction struct {
	ID         uint      `gorm:"primaryKey;autoIncrement"` // Уникальный идентификатор транзакции
	UserID     uint      `gorm:"index;not null"`           // Идентификатор пользователя, которому принадлежит транзакция
	Amount     float64   `gorm:"not null"`                 // Сумма транзакции (должна быть положительным числом)
	CategoryID uint      `gorm:"index;not null"`           // Идентификатор категории транзакции
	Date       time.Time `gorm:"not null"`                 // Дата проведения транзакции
	CreatedAt  time.Time // Время создания записи транзакции
	UpdatedAt  time.Time // Время последнего обновления записи транзакции
}

// NewTransaction создаёт новую транзакцию с заданными параметрами. Выполняется проверка корректности входных данных:
// - userID должен быть положительным числом,
// - categoryID должен быть положительным числом,
// - сумма транзакции должна быть положительной,
// - дата транзакции не может быть пустой.
// В случае некорректных данных возвращается ошибка валидации.
func NewTransaction(categoryID uint, userID uint, amount float64, date time.Time) (*Transaction, error) {
	if userID == 0 {
		return nil, errs.NewValidationError("ID пользователя должен быть положительным числом")
	}

	if categoryID == 0 {
		return nil, errs.NewValidationError("ID категории должен быть положительным числом")
	}

	if amount <= 0 {
		return nil, errs.NewValidationError("сумма транзакции должна быть положительным числом")
	}

	if date.IsZero() {
		return nil, errs.NewValidationError("дата транзакции не может быть пустой")
	}

	return &Transaction{
		CategoryID: categoryID,
		UserID:     userID,
		Date:       date,
		Amount:     amount,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}, nil
}
