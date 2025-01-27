package usecases

import (
	"time"

	"github.com/vladkaprelev/finance-go/internal/domain/model"
	"github.com/vladkaprelev/finance-go/internal/domain/repo"
	"github.com/vladkaprelev/finance-go/internal/errs"
)

type ITransactionUsecase interface {
	CreateTransaction(userID, categoryID, budgetID uint, amount float64, date time.Time) (*model.Transaction, error)
	UpdateTransaction(userID, categoryID, budgetID uint, amount float64, date time.Time) (*model.Transaction, error)
	DeleteTransaction(id uint) error

	FindTransactionsByUserID(userID uint, startDate time.Time, endDate time.Time) ([]*model.Transaction, error)
	FindTransactionsByCategoryID(categoryID uint, startDate time.Time, endDate time.Time) ([]*model.Transaction, error)
	FindTransactionsByBudgetID(budgetID uint, startDate time.Time, endDate time.Time) ([]*model.Transaction, error)
}

type transactionUsecase struct {
	transactionRepo repo.ITransactionRepository
}

func NewTransactionUsecase(transactionRepo repo.ITransactionRepository) ITransactionUsecase {
	return &transactionUsecase{
		transactionRepo: transactionRepo,
	}
}

func (u *transactionUsecase) CreateTransaction(
	userID, categoryID, budgetID uint,
	amount float64,
	date time.Time,
) (*model.Transaction, error) {
	if userID == 0 {
		return nil, errs.NewValidationError("CreateTransaction: некорректное значение userID")
	}

	if categoryID == 0 {
		return nil, errs.NewValidationError("CreateTransaction: некорректное значение categoryID")
	}

	if budgetID == 0 {
		return nil, errs.NewValidationError("CreateTransaction: некорректное значение budgetID")
	}

	if amount <= 0 {
		return nil, errs.NewValidationError("CreateTransaction: некорректное значение amount")
	}

	transaction := &model.Transaction{
		UserID:     userID,
		CategoryID: categoryID,
		BudgetID:   budgetID,
		Amount:     amount,
		Date:       date,
	}

	createdTransaction, err := u.transactionRepo.Create(transaction)
	if err != nil {
		return nil, err
	}

	return createdTransaction, nil
}

func (u *transactionUsecase) UpdateTransaction(
	userID, categoryID, budgetID uint,
	amount float64,
	date time.Time,
) (*model.Transaction, error) {
	if userID == 0 {
		return nil, errs.NewValidationError("CreateTransaction: некорректное значение userID")
	}

	if categoryID == 0 {
		return nil, errs.NewValidationError("CreateTransaction: некорректное значение categoryID")
	}

	if budgetID == 0 {
		return nil, errs.NewValidationError("CreateTransaction: некорректное значение budgetID")
	}

	if amount <= 0 {
		return nil, errs.NewValidationError("CreateTransaction: некорректное значение amount")
	}

	transaction := &model.Transaction{
		UserID:     userID,
		CategoryID: categoryID,
		BudgetID:   budgetID,
		Amount:     amount,
		Date:       date,
	}

	updatedTransaction, err := u.transactionRepo.Update(transaction)
	if err != nil {
		return nil, err
	}

	return updatedTransaction, nil
}

func (u *transactionUsecase) DeleteTransaction(id uint) error {
	err := u.transactionRepo.Delete(id)

	return err
}

func (u *transactionUsecase) FindTransactionsByUserID(
	userID uint,
	startDate time.Time,
	endDate time.Time,
) ([]*model.Transaction, error) {
	transactions, err := u.transactionRepo.FindByUserID(userID, startDate, endDate)

	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (u *transactionUsecase) FindTransactionsByCategoryID(
	userID uint,
	startDate time.Time,
	endDate time.Time,
) ([]*model.Transaction, error) {
	transactions, err := u.transactionRepo.FindByCategoryID(userID, startDate, endDate)

	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (u *transactionUsecase) FindTransactionsByBudgetID(
	userID uint,
	startDate time.Time,
	endDate time.Time,
) ([]*model.Transaction, error) {
	transactions, err := u.transactionRepo.FindByBudgetID(userID, startDate, endDate)

	if err != nil {
		return nil, err
	}

	return transactions, nil
}
