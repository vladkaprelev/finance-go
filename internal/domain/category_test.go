package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/vladkaprelev/finance-go/internal/errs"
)

func TestNewCategory(t *testing.T) {
	tests := []struct {
		name          string
		inputName     string
		inputUserID   uint
		inputType     CategoryType
		expectedError error
	}{
		{
			name:          "ValidCategory_Expense",
			inputName:     "Groceries",
			inputUserID:   1,
			inputType:     Expense,
			expectedError: nil,
		},
		{
			name:          "ValidCategory_Income",
			inputName:     "Salary",
			inputUserID:   2,
			inputType:     Income,
			expectedError: nil,
		},
		{
			name:          "EmptyName",
			inputName:     "",
			inputUserID:   1,
			inputType:     Expense,
			expectedError: errs.NewValidationError("название категории не может быть пустым"),
		},
		{
			name:          "InvalidCategoryType",
			inputName:     "Investment",
			inputUserID:   1,
			inputType:     CategoryType("invalid_type"),
			expectedError: errs.NewValidationError("тип категории некорректен"),
		},
		{
			name:          "ZeroUserID",
			inputName:     "Utilities",
			inputUserID:   0,
			inputType:     Expense,
			expectedError: errs.NewValidationError("ID пользователя должен быть положительным числом"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			category, err := NewCategory(tt.inputName, tt.inputUserID, tt.inputType)

			if tt.expectedError != nil {
				assert.Nil(t, category)
				assert.Equal(t, tt.expectedError.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, category)
				assert.Equal(t, tt.inputName, category.Name)
				assert.Equal(t, tt.inputUserID, category.UserID)
				assert.Equal(t, tt.inputType, category.Type)
				assert.WithinDuration(t, time.Now(), category.CreatedAt, time.Second)
				assert.WithinDuration(t, time.Now(), category.UpdatedAt, time.Second)
			}
		})
	}
}
