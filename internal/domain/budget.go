package domain

import (
	"time"

	"github.com/vladkaprelev/finance-go/internal/errs"
)

// Budget представляет бюджет, связанный с категорией расходов или доходов.
// Он содержит идентификаторы категории и пользователя, даты начала и окончания,
// а также временные метки создания и обновления записи.
type Budget struct {
	ID         uint `gorm:"primaryKey;autoIncrement"`
	CategotyID uint
	UserID     uint

	StartDate time.Time
	EndDate   time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewBudget конструктор создания budget
func NewBudget(
	categoryID uint,
	userID uint,
	startDate time.Time,
	endDate time.Time,
) (*Budget, error) {
	if userID == 0 {
		return nil, errs.NewValidationError("ID пользователя должен быть положительным числом")
	}

	if categoryID == 0 {
		return nil, errs.NewValidationError("ID категории должен быть положительным числом")
	}

	if endDate.Before(startDate) {
		return nil, errs.NewValidationError("Дата окончания не может быть раньше даты начала")
	}

	return &Budget{
		CategotyID: categoryID,
		UserID:     userID,
		StartDate:  startDate,
		EndDate:    endDate,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}, nil
}
