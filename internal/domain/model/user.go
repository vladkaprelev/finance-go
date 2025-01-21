package model

import (
	"regexp"
	"time"

	"github.com/vladkaprelev/finance-go/internal/errs"
	"golang.org/x/crypto/bcrypt"
)

const (
	MinPasswordLength = 8
)

// User — структура, представляющая пользователя.
type User struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Регулярное выражение для проверки наличия хотя бы одной буквы
var letterRegex = regexp.MustCompile(`[A-Za-z]`)

// Регулярное выражение для проверки email
var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

// Регулярное выражение для проверки наличия хотя бы одной цифры
var digitRegex = regexp.MustCompile(`\d`)

// ValidateEmail — метод для проверки корректности формата email.
func (u *User) ValidateEmail() error {
	if !emailRegex.MatchString(u.Email) {
		return errs.NewValidationError("некорректный формат email")
	}

	return nil
}

// ValidatePassword — метод для проверки корректности пароля (длина, наличие буквы и цифры).
func (u *User) ValidatePassword() error {
	// Проверка длины пароля
	if len(u.Password) < MinPasswordLength {
		return errs.NewValidationError("пароль должен содержать минимум 8 символов")
	}

	// Проверка наличия хотя бы одной буквы
	if !letterRegex.MatchString(u.Password) {
		return errs.NewValidationError("пароль должен содержать хотя бы одну букву")
	}

	// Проверка наличия хотя бы одной цифры
	if !digitRegex.MatchString(u.Password) {
		return errs.NewValidationError("пароль должен содержать хотя бы одну цифру")
	}

	return nil
}

// NewUser — функция для создания нового пользователя (с валидацией и хешированием пароля).
func NewUser(name, email, password string) (*User, error) {
	if name == "" {
		return nil, errs.NewValidationError("имя пользователя не может быть пустым")
	}

	if email == "" {
		return nil, errs.NewValidationError("email не может быть пустым")
	}

	if password == "" {
		return nil, errs.NewValidationError("password не может быть пустым")
	}

	user := &User{
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Валидация email и пароля
	if err := user.ValidateEmail(); err != nil {
		return nil, err
	}

	if err := user.ValidatePassword(); err != nil {
		return nil, err
	}

	// Хеширование пароля
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errs.NewValidationError("не удалось захешировать пароль")
	}

	user.Password = string(hashedPassword)

	return user, nil
}
