package usecases

import (
	"time"

	"github.com/vladkaprelev/finance-go/internal/domain/model"
	"github.com/vladkaprelev/finance-go/internal/domain/repo"
	"github.com/vladkaprelev/finance-go/internal/errs"
)

type IBudgetUsecase interface {
	CreateBudget(
		userID, categoryID uint,
		targetValue int32,
		startDate, endDate time.Time,
	) (*model.Budget, error)
	UpdateBudget(
		userID, categoryID uint,
		targetValue int32,
		startDate, endDate time.Time,
	) (*model.Budget, error)
	DeleteBudget(id uint) error

	FindBudgetsByUserID(userID uint, startDate time.Time, endDate time.Time) ([]*model.Budget, error)
	FindBudgetsByCategoryID(categoryID uint, startDate time.Time, endDate time.Time) ([]*model.Budget, error)
}

type budgetUsecase struct {
	budgetRepo repo.IBudgetRepository
}

func NewBudgetUseCase(budgetRepo repo.IBudgetRepository) IBudgetUsecase {
	return &budgetUsecase{
		budgetRepo: budgetRepo,
	}
}

func (u *budgetUsecase) CreateBudget(
	userID, categoryID uint,
	targetValue int32,
	startDate, endDate time.Time,
) (*model.Budget, error) {
	if userID == 0 {
		return nil, errs.NewValidationError("CreateBudget: некорректное значение userID")
	}

	if categoryID == 0 {
		return nil, errs.NewValidationError("CreateBudget: некорректное значение categoryID")
	}

	if targetValue <= 0 {
		return nil, errs.NewValidationError("CreateBudget: некорректное значение targetValue")
	}

	budget := &model.Budget{
		UserID:      userID,
		CategoryID:  categoryID,
		TargetValue: targetValue,
		StartDate:   startDate,
		EndDate:     endDate,
	}

	createdBudget, err := u.budgetRepo.Create(budget)
	if err != nil {
		return nil, err
	}

	return createdBudget, nil
}
func (u *budgetUsecase) UpdateBudget(
	userID, categoryID uint,
	targetValue int32,
	startDate, endDate time.Time,
) (*model.Budget, error) {
	if userID == 0 {
		return nil, errs.NewValidationError("UpdateBudget: некорректное значение userID")
	}

	if categoryID == 0 {
		return nil, errs.NewValidationError("UpdateBudget: некорректное значение categoryID")
	}

	if targetValue <= 0 {
		return nil, errs.NewValidationError("UpdateBudget: некорректное значение targetValue")
	}

	budget := &model.Budget{
		UserID:      userID,
		CategoryID:  categoryID,
		TargetValue: targetValue,
		StartDate:   startDate,
		EndDate:     endDate,
	}

	updatedBudget, err := u.budgetRepo.Update(budget)
	if err != nil {
		return nil, err
	}

	return updatedBudget, nil
}

func (u *budgetUsecase) DeleteBudget(id uint) error {
	err := u.budgetRepo.Delete(id)

	return err
}

func (u *budgetUsecase) FindBudgetsByUserID(
	userID uint,
	startDate time.Time,
	endDate time.Time,
) ([]*model.Budget, error) {
	budget, err := u.budgetRepo.FindByUserID(userID, startDate, endDate)

	if err != nil {
		return nil, err
	}

	return budget, nil
}

func (u *budgetUsecase) FindBudgetsByCategoryID(
	categoryID uint,
	startDate time.Time,
	endDate time.Time,
) ([]*model.Budget, error) {
	budget, err := u.budgetRepo.FindByCategoryID(categoryID, startDate, endDate)

	if err != nil {
		return nil, err
	}

	return budget, nil
}
