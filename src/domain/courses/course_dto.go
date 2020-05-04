package courses

import (
	"strings"

	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
)

type Course struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	CategoryID  int64  `json:"category_id"`
	Image       string `json:"image"`
	DateCreated string `json:"date_created"`
}

type Courses []Course

func (course *Course) Validate() *rest_errors.RestErr {
	course.Name = strings.TrimSpace(course.Name)
	if course.Name == "" {
		return rest_errors.NewBadRequestError("invalid course name")
	}
	if course.CategoryID <= 0 {
		return rest_errors.NewBadRequestError("invalid category id")
	}
	return nil
}
