package services

import (
	"github.com/dzikrisyafi/kursusvirtual_courses-api/src/domain/categories"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
)

var (
	CategoryService categoryServiceInterface = &categoryService{}
)

type categoryService struct{}

type categoryServiceInterface interface {
	CreateCategory(categories.Category) (*categories.Category, rest_errors.RestErr)
	GetCategory(int) (*categories.Category, rest_errors.RestErr)
	GetAllCategory() (categories.Categories, rest_errors.RestErr)
	UpdateCategory(categories.Category) (*categories.Category, rest_errors.RestErr)
	DeleteCategory(int) rest_errors.RestErr
}

func (s *categoryService) CreateCategory(category categories.Category) (*categories.Category, rest_errors.RestErr) {
	if err := category.Validate(); err != nil {
		return nil, err
	}

	if err := category.Save(); err != nil {
		return nil, err
	}

	return &category, nil
}

func (s *categoryService) GetCategory(categoryID int) (*categories.Category, rest_errors.RestErr) {
	dao := &categories.Category{ID: categoryID}
	if err := dao.Get(); err != nil {
		return nil, err
	}

	return dao, nil
}

func (s *categoryService) GetAllCategory() (categories.Categories, rest_errors.RestErr) {
	dao := &categories.Category{}
	return dao.GetAll()
}

func (s *categoryService) UpdateCategory(category categories.Category) (*categories.Category, rest_errors.RestErr) {
	current, err := s.GetCategory(category.ID)
	if err != nil {
		return nil, err
	}

	if err := category.Validate(); err != nil {
		return nil, err
	}

	current.Name = category.Name
	if err := current.Update(); err != nil {
		return nil, err
	}

	return current, nil
}

func (s *categoryService) DeleteCategory(categoryID int) rest_errors.RestErr {
	dao := &categories.Category{ID: categoryID}
	return dao.Delete()
}
