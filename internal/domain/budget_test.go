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
		categoryId    uint
		userId        uint
		startDate     time.Time
		endDate       time.Time
		expectedError error
	}{
		{
			name:          "ValidBudget",
			categoryId:    1,
			userId:        1,
			startDate:     time.Now(),
			endDate:       time.Now().AddDate(0, 1, 0),
			expectedError: nil,
		},
		{
			name:          "ZeroUserID",
			categoryId:    1,
			userId:        0,
			startDate:     time.Now(),
			endDate:       time.Now().AddDate(0, 1, 0),
			expectedError: errs.NewValidationError("ID пользователя должен быть положительным числом"),
		},
		{
			name:          "ZeroCategoryID",
			categoryId:    0,
			userId:        1,
			startDate:     time.Now(),
			endDate:       time.Now().AddDate(0, 1, 0),
			expectedError: errs.NewValidationError("ID категории должен быть положительным числом"),
		},
		{
			name:          "EndDateBeforeStartDate",
			categoryId:    1,
			userId:        1,
			startDate:     time.Now(),
			endDate:       time.Now().AddDate(0, -1, 0),
			expectedError: errs.NewValidationError("Дата окончания не может быть раньше даты начала"),
		},
		{
			name:          "StartDateEqualsEndDate",
			categoryId:    1,
			userId:        1,
			startDate:     time.Now(),
			endDate:       time.Now(),
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			budget, err := NewBudget(tt.categoryId, tt.userId, tt.startDate, tt.endDate)

			if tt.expectedError != nil {
				assert.Nil(t, budget)
				assert.Equal(t, tt.expectedError.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, budget)
				assert.Equal(t, tt.categoryId, budget.CategotyID)
				assert.Equal(t, tt.userId, budget.UserID)
				assert.Equal(t, tt.startDate, budget.StartDate)
				assert.Equal(t, tt.endDate, budget.EndDate)
				assert.WithinDuration(t, time.Now(), budget.CreatedAt, time.Second)
				assert.WithinDuration(t, time.Now(), budget.UpdatedAt, time.Second)
			}
		})
	}
}
