package usecases

import (
	"github.com/vladkaprelev/finance-go/internal/domain/model"
	"github.com/vladkaprelev/finance-go/internal/domain/repo"
	"github.com/vladkaprelev/finance-go/internal/errs"
)

type ICategoryUsecase interface {
	CreateCategory(name string, userID uint, categoryType model.CategoryType) (*model.Category, error)
	UpdateCategory(name string, userID uint, categoryType model.CategoryType) (*model.Category, error)
	GetCategoryByID(id uint) (*model.Category, error)
	GetCategoriesByUserID(userID uint) ([]*model.Category, error)
	DeleteCategoryByID(id uint) error
}

type categoryUsecase struct {
	categoryRepo repo.ICategoryRepository
}

func NewCategoryUseCase(repo repo.ICategoryRepository) ICategoryUsecase {
	return &categoryUsecase{
		categoryRepo: repo,
	}
}

func (u *categoryUsecase) CreateCategory(
	name string,
	userID uint,
	categoryType model.CategoryType,
) (*model.Category, error) {
	if name == "" {
		return nil, errs.NewValidationError("CreateCategory: некорректное название категории")
	}

	if userID == 0 {
		return nil, errs.NewValidationError("CreateCategory: некорректное значение userID")
	}

	if !categoryType.IsValid() {
		return nil, errs.NewValidationError("CreateCategory: некорректный тип категории")
	}

	category := &model.Category{
		Name:   name,
		UserID: userID,
		Type:   categoryType,
	}

	createdCategory, err := u.categoryRepo.Create(category)
	if err != nil {
		return nil, err
	}

	return createdCategory, nil
}

func (u *categoryUsecase) UpdateCategory(
	name string,
	userID uint,
	categoryType model.CategoryType,
) (*model.Category, error) {
	if name == "" {
		return nil, errs.NewValidationError("CreateCategory: некорректное название категории")
	}

	if userID == 0 {
		return nil, errs.NewValidationError("CreateCategory: некорректное значение userID")
	}

	if !categoryType.IsValid() {
		return nil, errs.NewValidationError("CreateCategory: некорректный тип категории")
	}

	category := &model.Category{
		Name:   name,
		UserID: userID,
		Type:   categoryType,
	}

	updatedCategory, err := u.categoryRepo.Update(category)
	if err != nil {
		return nil, err
	}

	return updatedCategory, nil
}

func (u *categoryUsecase) GetCategoryByID(id uint) (*model.Category, error) {
	category, err := u.categoryRepo.FindByID(id)

	if err != nil {
		return nil, err
	}

	return category, nil
}

func (u *categoryUsecase) DeleteCategoryByID(id uint) error {
	err := u.categoryRepo.Delete(id)

	return err
}

func (u *categoryUsecase) GetCategoriesByUserID(userID uint) ([]*model.Category, error) {
	categories, err := u.categoryRepo.FindByUserID(userID)

	if err != nil {
		return nil, err
	}

	return categories, nil
}
