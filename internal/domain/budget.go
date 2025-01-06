package domain

import (
	"time"

	"github.com/vladkaprelev/finance-go/internal/errs"
)

type Budget struct {
	ID         uint `gorm:"primaryKey;autoIncrement"`
	CategotyID uint
	UserID     uint

	StartDate time.Time
	EndDate   time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewBudget(
	categoryId uint,
	userId uint,
	startDate time.Time,
	endDate time.Time,
) (*Budget, error) {

	if userId == 0 {
		return nil, errs.NewValidationError("ID пользователя должен быть положительным числом")
	}

	if categoryId == 0 {
		return nil, errs.NewValidationError("ID категории должен быть положительным числом")
	}

	if endDate.Before(startDate) {
		return nil, errs.NewValidationError("Дата окончания не может быть раньше даты начала")
	}

	return &Budget{
		CategotyID: categoryId,
		UserID:     userId,
		StartDate:  startDate,
		EndDate:    endDate,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}, nil
}
