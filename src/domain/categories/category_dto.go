package categories

import (
	"html"
	"strings"

	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
)

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Categories []Category

func (category *Category) Validate() rest_errors.RestErr {
	category.Name = html.EscapeString(strings.TrimSpace(category.Name))
	if category.Name == "" {
		return rest_errors.NewBadRequestError("invalid category name")
	}

	return nil
}
