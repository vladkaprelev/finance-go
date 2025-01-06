package domain

import (
	"time"

	"github.com/vladkaprelev/finance-go/internal/errs"
)

type CategoryType string

const (
	Expense CategoryType = "expense" // Расход
	Income  CategoryType = "income"  // Доход
)

var ValidCategoryTypes = []CategoryType{
	Expense,
	Income,
}

func (ct CategoryType) IsValid() bool {
	for _, validType := range ValidCategoryTypes {
		if ct == validType {
			return true
		}
	}
	return false
}

type Category struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	Name      string
	UserID    uint
	Type      CategoryType
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewCategory(name string, userId uint, categoryType CategoryType) (*Category, error) {

	if name == "" {
		return nil, errs.NewValidationError("название категории не может быть пустым")
	}

	if !categoryType.IsValid() {
		return nil, errs.NewValidationError("тип категории некорректен")
	}

	if userId == 0 {
		return nil, errs.NewValidationError("ID пользователя должен быть положительным числом")
	}

	return &Category{
		Name:      name,
		UserID:    userId,
		Type:      categoryType,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
