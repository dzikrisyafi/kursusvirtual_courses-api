package services

import (
	"github.com/dzikrisyafi/kursusvirtual_courses-api/src/domain/enrolls"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
)

var (
	EnrollsService enrollsServiceInterface = &enrollsService{}
)

type enrollsService struct{}

type enrollsServiceInterface interface {
	GetCourseByUserID(int) (*enrolls.User, rest_errors.RestErr)
	CreateEnroll(enrolls.Enroll) (*enrolls.Enroll, rest_errors.RestErr)
	DeleteEnroll(int) rest_errors.RestErr
}

func (s *enrollsService) GetCourseByUserID(userID int) (*enrolls.User, rest_errors.RestErr) {
	result := &enrolls.User{UserID: userID}
	if err := result.GetCourseByUserID(); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *enrollsService) CreateEnroll(req enrolls.Enroll) (*enrolls.Enroll, rest_errors.RestErr) {
	dao := &enrolls.Enroll{
		UserID:   req.UserID,
		CourseID: req.CourseID,
		Cohort:   req.Cohort,
	}

	if err := dao.Save(); err != nil {
		return nil, err
	}

	return dao, nil
}

func (s *enrollsService) DeleteEnroll(enrollID int) rest_errors.RestErr {
	dao := enrolls.Enroll{ID: enrollID}
	return dao.Delete()
}
