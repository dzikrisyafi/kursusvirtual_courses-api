package courses

import (
	"strings"

	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
	"golang.org/x/net/html"
)

type Course struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	CategoryID  int    `json:"category_id"`
	Image       string `json:"image"`
	DateCreated string `json:"date_created"`
}

type Courses []Course

func (course *Course) Validate() rest_errors.RestErr {
	course.Name = html.EscapeString(strings.TrimSpace(course.Name))
	if course.Name == "" {
		return rest_errors.NewBadRequestError("invalid course name")
	}

	if course.CategoryID <= 0 {
		return rest_errors.NewBadRequestError("invalid category id")
	}

	return nil
}
