package enrolls

import (
	"fmt"
	"time"

	"errors"

	"github.com/dzikrisyafi/kursusvirtual_courses-api/src/datasources/mysql/courses_db"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/logger"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
	"github.com/golang-restclient/rest"
)

var (
	oauthRestClient = rest.RequestBuilder{
		BaseURL: "http://localhost:8000",
		Timeout: 100 * time.Millisecond,
	}
)

const (
	queryInsertEnroll      = `INSERT INTO enrolls (user_id, course_id, cohort_id) VALUES (?, ?, ?);`
	queryGetCourseByUserID = `SELECT course_id, name, image FROM enrolls INNER JOIN courses ON course_id=courses.id WHERE user_id=?;`
)

func (user *User) GetCourseByUserID() rest_errors.RestErr {
	stmt, err := courses_db.DbConn().Prepare(queryGetCourseByUserID)
	if err != nil {
		logger.Error("error when trying to get course by user id", err)
		return rest_errors.NewInternalServerError("error when trying to get course", errors.New("database error"))
	}
	defer stmt.Close()

	rows, err := stmt.Query(user.UserID)
	if err != nil {
		logger.Error("error when trying to get course by user id", err)
		return rest_errors.NewInternalServerError("error when trying to get course", errors.New("database error"))
	}
	defer rows.Close()

	var course UserCourse
	for rows.Next() {
		if err := rows.Scan(&course.CourseID, &course.CourseName, &course.Image); err != nil {
			logger.Error("error when trying to scan course rows into course struct", err)
			return rest_errors.NewInternalServerError("error when trying to get course", errors.New("database error"))
		}

		user.Courses = append(user.Courses, course)
	}

	if len(user.Courses) == 0 {
		return rest_errors.NewNotFoundError(fmt.Sprintf("no courses matching given user id %d", user.UserID))
	}

	return nil
}

func (enroll *Enroll) Save() rest_errors.RestErr {
	stmt, err := courses_db.DbConn().Prepare(queryInsertEnroll)
	if err != nil {
		logger.Error("error when trying to prepare save enroll statement", err)
		return rest_errors.NewInternalServerError("error when trying to save enroll", errors.New("database error"))
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(enroll.UserID, enroll.CourseID, enroll.Cohort)
	if saveErr != nil {
		logger.Error("error when trying to save enroll", saveErr)
		return rest_errors.NewInternalServerError("error when trying to save enroll", errors.New("database error"))
	}

	enrollID, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when tryin to get last insert id after creating a new enroll", err)
		return rest_errors.NewInternalServerError("error when trying to save enroll", errors.New("database error"))
	}
	enroll.ID = enrollID

	return nil
}

func (enroll *Enroll) EnrollRequest() rest_errors.RestErr {
	response := oauthRestClient.Post("/internal/enrolls", enroll)

	if response == nil || response.Response == nil {
		return rest_errors.NewInternalServerError("invalid rest client response when trying to get access token", errors.New("rest client error"))
	}

	if response.StatusCode > 299 {
		apiErr, err := rest_errors.NewRestErrorFromBytes(response.Bytes())
		if err != nil {
			return rest_errors.NewInternalServerError("invalid error interface when trying to get access token", err)
		}
		return apiErr
	}
	return nil
}
