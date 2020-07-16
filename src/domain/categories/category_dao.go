package categories

import (
	"errors"
	"strings"

	"github.com/dzikrisyafi/kursusvirtual_courses-api/src/datasources/mysql/courses_db"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/logger"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
)

const (
	queryInsertCategory = `INSERT INTO categories (name) VALUES (?);`
	queryGetCategory    = `SELECT name FROM categories WHERE id=?;`
	queryGetAllCategory = `SELECT id, name FROM categories;`
	queryUpdateCategory = `UPDATE categories SET name=? WHERE id=?;`
	queryDeleteCategory = `DELETE FROM categories WHERE id=?;`
)

func (category *Category) Save() rest_errors.RestErr {
	stmt, err := courses_db.DbConn().Prepare(queryInsertCategory)
	if err != nil {
		logger.Error("error when trying to prepare save category statement", err)
		return rest_errors.NewInternalServerError("error when trying to save category", errors.New("database error"))
	}
	defer stmt.Close()

	result, err := stmt.Exec(category.Name)
	if err != nil {
		logger.Error("error when trying to save category", err)
		return rest_errors.NewInternalServerError("error when trying to save category", errors.New("database error"))
	}

	categoryID, err := result.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get last insert id after creating a new category", err)
		return rest_errors.NewInternalServerError("error when trying to save category", errors.New("database error"))
	}
	category.ID = int(categoryID)

	return nil
}

func (category *Category) Get() rest_errors.RestErr {
	stmt, err := courses_db.DbConn().Prepare(queryGetCategory)
	if err != nil {
		logger.Error("error when trying to prepare get category by id statement", err)
		return rest_errors.NewInternalServerError("error when trying to get category", errors.New("database error"))
	}
	defer stmt.Close()

	row := stmt.QueryRow(category.ID)
	if getErr := row.Scan(&category.Name); getErr != nil {
		if strings.Contains(getErr.Error(), "no rows") {
			return rest_errors.NewNotFoundError("no category matching with given id")
		}
		logger.Error("error when trying to get category by id", err)
		return rest_errors.NewInternalServerError("error when trying to get category", errors.New("database error"))
	}

	return nil
}

func (category *Category) GetAll() ([]Category, rest_errors.RestErr) {
	stmt, err := courses_db.DbConn().Prepare(queryGetAllCategory)
	if err != nil {
		logger.Error("error when trying to prepare get all category statement", err)
		return nil, rest_errors.NewInternalServerError("error when trying to get category", errors.New("database error"))
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		logger.Error("error when trying to get all category", err)
		return nil, rest_errors.NewInternalServerError("error when trying to get all category", errors.New("database error"))
	}
	defer rows.Close()

	result := make([]Category, 0)
	for rows.Next() {
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			logger.Error("error when trying to scan category row into category struct", err)
			return nil, rest_errors.NewInternalServerError("error when trying to get all category", errors.New("database error"))
		}

		result = append(result, *category)
	}

	if len(result) == 0 {
		return nil, rest_errors.NewNotFoundError("no category rows in result set")
	}

	return result, nil
}

func (category *Category) Update() rest_errors.RestErr {
	stmt, err := courses_db.DbConn().Prepare(queryUpdateCategory)
	if err != nil {
		logger.Error("error when trying to prepare update category by id statement", err)
		return rest_errors.NewInternalServerError("error when trying to update category", errors.New("database error"))
	}
	defer stmt.Close()

	_, err = stmt.Exec(category.Name, category.ID)
	if err != nil {
		logger.Error("error when trying to update category", err)
		return rest_errors.NewInternalServerError("error when trying to update category", errors.New("database error"))
	}

	return nil
}

func (category *Category) Delete() rest_errors.RestErr {
	stmt, err := courses_db.DbConn().Prepare(queryDeleteCategory)
	if err != nil {
		logger.Error("error when trying prepare delete category by id statement", err)
		return rest_errors.NewInternalServerError("error when trying to delete category", errors.New("database error"))
	}
	defer stmt.Close()

	if _, err = stmt.Exec(category.ID); err != nil {
		logger.Error("error when trying to delete category by id", err)
		return rest_errors.NewInternalServerError("error when trying to delete category", errors.New("database error"))
	}

	return nil
}
