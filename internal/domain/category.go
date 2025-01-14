package domain

import (
	"time"

	"github.com/vladkaprelev/finance-go/internal/errs"
)

// CategoryType представляет тип категории (например, "expense" для расходов или "income" для доходов).
type CategoryType string

const (
	// Expense означает категорию расходов.
	Expense CategoryType = "expense"
	// Income означает категорию доходов.
	Income CategoryType = "income"
)

// ValidCategoryTypes содержит список допустимых типов категории.
var ValidCategoryTypes = []CategoryType{
	Expense,
	Income,
}

// IsValid проверяет, является ли данный тип категории допустимым.
func (ct CategoryType) IsValid() bool {
	for _, validType := range ValidCategoryTypes {
		if ct == validType {
			return true
		}
	}

	return false
}

// Category представляет категорию транзакций с указанием имени, пользователя и типа категории.
type Category struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	Name      string
	UserID    uint
	Type      CategoryType
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewCategory создаёт новую категорию с заданными параметрами, выполняя валидацию входных данных.
// Возвращает ошибку валидации, если входные данные некорректны.
func NewCategory(name string, userID uint, categoryType CategoryType) (*Category, error) {
	if name == "" {
		return nil, errs.NewValidationError("название категории не может быть пустым")
	}

	if !categoryType.IsValid() {
		return nil, errs.NewValidationError("тип категории некорректен")
	}

	if userID == 0 {
		return nil, errs.NewValidationError("ID пользователя должен быть положительным числом")
	}

	return &Category{
		Name:      name,
		UserID:    userID,
		Type:      categoryType,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
