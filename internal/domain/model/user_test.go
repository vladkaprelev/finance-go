package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/vladkaprelev/finance-go/internal/errs"
	"golang.org/x/crypto/bcrypt"
)

func TestNewUser(t *testing.T) {
	tests := []struct {
		name          string
		inputEmail    string
		inputPassword string
		expectedError error
	}{
		{
			name:          "ValidUser",
			inputEmail:    "ivan.ivanov@example.com",
			inputPassword: "Password123",
			expectedError: nil,
		},
		{
			name:          "EmptyEmail",
			inputEmail:    "",
			inputPassword: "Password123",
			expectedError: errs.NewValidationError("email не может быть пустым"),
		},
		{
			name:          "EmptyPassword",
			inputEmail:    "ivan.ivanov@example.com",
			inputPassword: "",
			expectedError: errs.NewValidationError("password не может быть пустым"),
		},
		{
			name:          "InvalidEmailFormat",
			inputEmail:    "ivan.ivanov@invalid", // Некорректный формат email
			inputPassword: "Password123",
			expectedError: errs.NewValidationError("некорректный формат email"),
		},
		{
			name:          "PasswordTooShort",
			inputEmail:    "ivan.ivanov@example.com",
			inputPassword: "Pass1", // Менее 8 символов
			expectedError: errs.NewValidationError("пароль должен содержать минимум 8 символов"),
		},
		{
			name:          "PasswordNoLetters",
			inputEmail:    "ivan.ivanov@example.com",
			inputPassword: "12345678", // Нет букв
			expectedError: errs.NewValidationError("пароль должен содержать хотя бы одну букву"),
		},
		{
			name:          "PasswordNoDigits",
			inputEmail:    "ivan.ivanov@example.com",
			inputPassword: "Password", // Нет цифр
			expectedError: errs.NewValidationError("пароль должен содержать хотя бы одну цифру"),
		},
		{
			name:          "PasswordHashingError",
			inputEmail:    "ivan.ivanov@example.com",
			inputPassword: "Password123",
			// Для симуляции ошибки хеширования можно использовать мокирование, но это выходит за рамки простого теста.
			// Поэтому данный тест не реализован.
			// В реальном проекте можно использовать интерфейсы для абстрагирования bcrypt и мокировать их.
			expectedError: nil, // Placeholder
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := NewUser(tt.inputEmail, tt.inputPassword)

			if tt.expectedError != nil {
				assert.Nil(t, user)
				assert.NotNil(t, err)
				assert.Equal(t, tt.expectedError.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, user)
				assert.Equal(t, tt.inputEmail, user.Email)
				// Проверяем, что пароль захеширован и отличается от исходного
				assert.NotEqual(t, tt.inputPassword, user.Password)
				// Проверяем, что захешированный пароль совпадает с исходным
				errHash := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(tt.inputPassword))
				assert.NoError(t, errHash)

				// Проверяем временные метки
				assert.WithinDuration(t, time.Now(), user.CreatedAt, time.Second)
				assert.WithinDuration(t, time.Now(), user.UpdatedAt, time.Second)
			}
		})
	}
}
