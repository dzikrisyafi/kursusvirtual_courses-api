package courses

import (
	"errors"
	"fmt"

	"github.com/dzikrisyafi/kursusvirtual_courses-api/src/datasources/mysql/courses_db"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/logger"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
)

const (
	queryGetCourse        = `SELECT id, name, category_id, image, date_created FROM courses WHERE id=?;`
	queryInsertCourse     = `INSERT INTO courses (name, category_id, image, date_created) VALUES (?, ?, ?, ?);`
	queryUpdateCourse     = `UPDATE courses SET name=?, category_id=?, image=? WHERE id=?;`
	queryDeleteCourse     = `DELETE FROM courses WHERE id=?;`
	queryFindCourseByName = `SELECT id, name, category_id, image, date_created FROM courses WHERE name LIKE ?`
)

func (course *Course) Save() rest_errors.RestErr {
	stmt, err := courses_db.DbConn().Prepare(queryInsertCourse)
	if err != nil {
		logger.Error("error when trying to prepare save course statement", err)
		return rest_errors.NewInternalServerError("error when trying to save course", errors.New("database error"))
	}
	defer stmt.Close()

	result, err := stmt.Exec(course.Name, course.CategoryID, course.Image, course.DateCreated)
	if err != nil {
		logger.Error("error when trying to save course", err)
		return rest_errors.NewInternalServerError("error when trying to save course", errors.New("database error"))
	}

	courseID, err := result.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get last insert id after creating a new course", err)
		return rest_errors.NewInternalServerError("error when trying to save course", errors.New("database error"))
	}
	course.ID = int(courseID)

	return nil
}

func (course *Course) Get() rest_errors.RestErr {
	stmt, err := courses_db.DbConn().Prepare(queryGetCourse)
	if err != nil {
		logger.Error("error when trying to prepare get course by id statement", err)
		return rest_errors.NewInternalServerError("error when trying to get course", errors.New("database error"))
	}
	defer stmt.Close()

	result := stmt.QueryRow(course.ID)
	if err = result.Scan(&course.ID, &course.Name, &course.CategoryID, &course.Image, &course.DateCreated); err != nil {
		logger.Error("error when trying to get course by id", err)
		return rest_errors.NewInternalServerError("error when trying to get course", errors.New("database error"))
	}

	return nil
}

func (course *Course) Update() rest_errors.RestErr {
	stmt, err := courses_db.DbConn().Prepare(queryUpdateCourse)
	if err != nil {
		logger.Error("error when trying to prepare update course statement", err)
		return rest_errors.NewInternalServerError("error when trying to update course", errors.New("database error"))
	}
	defer stmt.Close()

	_, err = stmt.Exec(course.Name, course.CategoryID, course.Image, course.ID)
	if err != nil {
		logger.Error("error when trying to update course", err)
		return rest_errors.NewInternalServerError("error when trying to update course", errors.New("database error"))
	}

	return nil
}

func (course *Course) Delete() rest_errors.RestErr {
	stmt, err := courses_db.DbConn().Prepare(queryDeleteCourse)
	if err != nil {
		logger.Error("error when trying to prepare delete course by id statement", err)
		return rest_errors.NewInternalServerError("error when trying to delete course", errors.New("database error"))
	}
	defer stmt.Close()

	if _, err = stmt.Exec(course.ID); err != nil {
		logger.Error("error when trying to delete course by id", err)
		return rest_errors.NewInternalServerError("error when trying to delete course", errors.New("database error"))
	}

	return nil
}

func (course *Course) FindCourseByName() ([]Course, rest_errors.RestErr) {
	stmt, err := courses_db.DbConn().Prepare(queryFindCourseByName)
	if err != nil {
		logger.Error("error when trying to prepare find course by name statement", err)
		return nil, rest_errors.NewInternalServerError("error when trying to find course", errors.New("database error"))
	}
	defer stmt.Close()

	rows, err := stmt.Query(course.Name + "%")
	if err != nil {
		logger.Error("error when trying to find course by name", err)
		return nil, rest_errors.NewInternalServerError("error when trying to find course", errors.New("database error"))
	}
	defer rows.Close()

	result := make([]Course, 0)
	for rows.Next() {
		if err := rows.Scan(&course.ID, &course.Name, &course.CategoryID, &course.Image, &course.DateCreated); err != nil {
			logger.Error("error when trying to scan course row into course struct", err)
			return nil, rest_errors.NewInternalServerError("error when trying to find course", errors.New("database error"))
		}

		result = append(result, *course)
	}

	if len(result) == 0 {
		return nil, rest_errors.NewNotFoundError(fmt.Sprintf("no course matching name %s", course.Name))
	}

	return result, nil
}
