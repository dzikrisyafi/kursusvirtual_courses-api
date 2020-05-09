package enrolls

import "github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"

type User struct {
	UserID  int64        `json:"user_id"`
	Courses []UserCourse `json:"courses"`
}

type UserCourse struct {
	CourseID   int64  `json:"course_id"`
	CourseName string `json:"name"`
	Image      string `json:"image"`
}

type Enroll struct {
	ID       int64  `json:"id"`
	UserID   int64  `json:"user_id"`
	CourseID int64  `json:"course_id"`
	Cohort   string `json:"cohort"`
}

type EnrollRequest struct {
	UserID   int64  `json:"user_id"`
	CourseID int64  `json:"course_id"`
	Cohort   string `json:"cohort"`
}

func (enroll *Enroll) Validate() rest_errors.RestErr {
	if enroll.UserID <= 0 {
		return rest_errors.NewBadRequestError("invalid user id")
	}

	if enroll.CourseID <= 0 {
		return rest_errors.NewBadRequestError("invalid course id")
	}

	return nil
}
