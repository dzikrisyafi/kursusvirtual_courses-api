package enrolls

import (
	"fmt"

	"errors"

	"github.com/dzikrisyafi/kursusvirtual_courses-api/src/datasources/mysql/courses_db"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/logger"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
)

const (
	queryInsertEnroll      = `INSERT INTO enrolls (user_id, course_id, cohort_id) VALUES (?, ?, ?);`
	queryGetCourseByUserID = `SELECT course_id, courses.name, image, cohort.name FROM enrolls INNER JOIN courses ON course_id=courses.id INNER JOIN cohort ON cohort_id=cohort.id WHERE user_id=?;`
	queryDeleteEnroll      = `DELETE FROM enrolls WHERE id=?;`
)

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
	enroll.ID = int(enrollID)

	return nil
}

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
		if err := rows.Scan(&course.CourseID, &course.CourseName, &course.Image, &course.Cohort); err != nil {
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

func (enroll *Enroll) Delete() rest_errors.RestErr {
	stmt, err := courses_db.DbConn().Prepare(queryDeleteEnroll)
	if err != nil {
		logger.Error("error when trying to prepare delete enroll by id statement", err)
		return rest_errors.NewInternalServerError("error when trying to delete enroll", errors.New("database error"))
	}
	defer stmt.Close()

	if _, err = stmt.Exec(enroll.ID); err != nil {
		logger.Error("error when trying to delete enroll by id", err)
		return rest_errors.NewInternalServerError("error when trying to delete enroll", errors.New("database error"))
	}

	return nil
}
