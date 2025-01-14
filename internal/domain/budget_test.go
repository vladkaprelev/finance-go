package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/vladkaprelev/finance-go/internal/errs"
)

func TestNewBudget(t *testing.T) {
	tests := []struct {
		name          string
		categoryID    uint
		userID        uint
		startDate     time.Time
		endDate       time.Time
		expectedError error
	}{
		{
			name:          "ValidBudget",
			categoryID:    1,
			userID:        1,
			startDate:     time.Now(),
			endDate:       time.Now().AddDate(0, 1, 0),
			expectedError: nil,
		},
		{
			name:          "ZeroUserID",
			categoryID:    1,
			userID:        0,
			startDate:     time.Now(),
			endDate:       time.Now().AddDate(0, 1, 0),
			expectedError: errs.NewValidationError("ID пользователя должен быть положительным числом"),
		},
		{
			name:          "ZeroCategoryID",
			categoryID:    0,
			userID:        1,
			startDate:     time.Now(),
			endDate:       time.Now().AddDate(0, 1, 0),
			expectedError: errs.NewValidationError("ID категории должен быть положительным числом"),
		},
		{
			name:          "EndDateBeforeStartDate",
			categoryID:    1,
			userID:        1,
			startDate:     time.Now(),
			endDate:       time.Now().AddDate(0, -1, 0),
			expectedError: errs.NewValidationError("Дата окончания не может быть раньше даты начала"),
		},
		{
			name:          "StartDateEqualsEndDate",
			categoryID:    1,
			userID:        1,
			startDate:     time.Now(),
			endDate:       time.Now(),
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			budget, err := NewBudget(tt.categoryID, tt.userID, tt.startDate, tt.endDate)

			if tt.expectedError != nil {
				assert.Nil(t, budget)
				assert.Equal(t, tt.expectedError.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, budget)
				assert.Equal(t, tt.categoryID, budget.CategotyID)
				assert.Equal(t, tt.userID, budget.UserID)
				assert.Equal(t, tt.startDate, budget.StartDate)
				assert.Equal(t, tt.endDate, budget.EndDate)
				assert.WithinDuration(t, time.Now(), budget.CreatedAt, time.Second)
				assert.WithinDuration(t, time.Now(), budget.UpdatedAt, time.Second)
			}
		})
	}
}
