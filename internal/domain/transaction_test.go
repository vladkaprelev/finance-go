// internal/domain/transaction_test.go
package domain

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/vladkaprelev/finance-go/internal/errs"
)

func TestNewTransaction(t *testing.T) {
	// Определяем тестовые случаи
	tests := []struct {
		name        string
		userID      uint
		categoryID  uint
		amount      float64
		date        time.Time
		expectError bool
		errorMsg    string
	}{
		{
			name:        "ValidTransaction",
			userID:      1,
			categoryID:  2,
			amount:      100.50,
			date:        time.Now(),
			expectError: false,
		},
		{
			name:        "ZeroUserID",
			userID:      0,
			categoryID:  2,
			amount:      100.50,
			date:        time.Now(),
			expectError: true,
			errorMsg:    "ID пользователя должен быть положительным числом",
		},
		{
			name:        "ZeroCategoryID",
			userID:      1,
			categoryID:  0,
			amount:      100.50,
			date:        time.Now(),
			expectError: true,
			errorMsg:    "ID категории должен быть положительным числом",
		},
		{
			name:        "NegativeAmount",
			userID:      1,
			categoryID:  2,
			amount:      -50.00,
			date:        time.Now(),
			expectError: true,
			errorMsg:    "сумма транзакции должна быть положительным числом",
		},
		{
			name:        "ZeroDate",
			userID:      1,
			categoryID:  2,
			amount:      100.50,
			date:        time.Time{},
			expectError: true,
			errorMsg:    "дата транзакции не может быть пустой",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transaction, err := NewTransaction(tt.categoryID, tt.userID, tt.amount, tt.date)

			if tt.expectError {
				assert.Nil(t, transaction, "Expected transaction to be nil on error")
				assert.Error(t, err, "Expected an error but got none")

				var appErr *errs.AppError
				if errors.As(err, &appErr) {
					assert.Equal(t, tt.errorMsg, appErr.Message, "Error message does not match")
				} else {
					t.Fatalf("Expected error to be of type *errs.AppError, got %T", err)
				}
			} else {
				assert.NoError(t, err, "Did not expect an error but got one")
				assert.NotNil(t, transaction, "Expected transaction to be created")

				assert.Equal(t, tt.userID, transaction.UserID, "UserID does not match")
				assert.Equal(t, tt.categoryID, transaction.CategoryID, "CategoryID does not match")
				assert.Equal(t, tt.amount, transaction.Amount, "Amount does not match")
				assert.Equal(t, tt.date, transaction.Date, "Date does not match")

				// Проверка временных меток
				assert.WithinDuration(t, time.Now(), transaction.CreatedAt, time.Second, "CreatedAt timestamp is not recent")
				assert.WithinDuration(t, time.Now(), transaction.UpdatedAt, time.Second, "UpdatedAt timestamp is not recent")
			}
		})
	}
}
