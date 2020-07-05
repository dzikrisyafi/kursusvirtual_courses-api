package services

import (
	"github.com/dzikrisyafi/kursusvirtual_courses-api/src/domain/courses"
	"github.com/dzikrisyafi/kursusvirtual_courses-api/src/repository/rest"
	"github.com/dzikrisyafi/kursusvirtual_courses-api/src/utils/date_utils"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
)

var (
	CoursesService coursesServiceInterface = &coursesService{}
)

type coursesService struct {
}

type coursesServiceInterface interface {
	CreateCourse(courses.Course) (*courses.Course, rest_errors.RestErr)
	GetCourse(int) (*courses.Course, rest_errors.RestErr)
	GetAllCourse() (courses.Courses, rest_errors.RestErr)
	UpdateCourse(bool, courses.Course) (*courses.Course, rest_errors.RestErr)
	DeleteCourse(int, string) rest_errors.RestErr
	SearchCourse(string) (courses.Courses, rest_errors.RestErr)
}

func (s *coursesService) CreateCourse(course courses.Course) (*courses.Course, rest_errors.RestErr) {
	if err := course.Validate(); err != nil {
		return nil, err
	}

	course.Image = "course.jpg"
	course.DateCreated = date_utils.GetNowDBFormat()
	if err := course.Save(); err != nil {
		return nil, err
	}

	return &course, nil
}

func (s *coursesService) GetCourse(courseID int) (*courses.Course, rest_errors.RestErr) {
	result := &courses.Course{ID: courseID}
	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *coursesService) GetAllCourse() (courses.Courses, rest_errors.RestErr) {
	dao := &courses.Course{}
	return dao.GetAll()
}

func (s *coursesService) UpdateCourse(isPartial bool, course courses.Course) (*courses.Course, rest_errors.RestErr) {
	current, err := s.GetCourse(course.ID)
	if err != nil {
		return nil, err
	}

	if isPartial {
		if course.Name != "" {
			current.Name = course.Name
		}

		if course.CategoryID != 0 {
			current.CategoryID = course.CategoryID
		}

	} else {
		if err := course.Validate(); err != nil {
			return nil, err
		}

		current.Name = course.Name
		current.CategoryID = course.CategoryID
	}

	if err := current.Update(); err != nil {
		return nil, err
	}

	return current, nil
}

func (s *coursesService) DeleteCourse(courseID int, at string) rest_errors.RestErr {
	dao := courses.Course{ID: courseID}

	if err := rest.TopicsRepository.DeleteTopics(courseID, at); err != nil {
		return err
	}

	return dao.Delete()
}

func (s *coursesService) SearchCourse(courseName string) (courses.Courses, rest_errors.RestErr) {
	dao := courses.Course{Name: courseName}
	return dao.FindCourseByName()
}
