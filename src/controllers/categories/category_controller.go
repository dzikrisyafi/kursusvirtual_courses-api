package categories

import (
	"net/http"

	"github.com/dzikrisyafi/kursusvirtual_courses-api/src/domain/categories"
	"github.com/dzikrisyafi/kursusvirtual_courses-api/src/services"
	"github.com/dzikrisyafi/kursusvirtual_oauth-go/oauth"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/controller_utils"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_resp"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var category categories.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	result, err := services.CategoryService.CreateCategory(category)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	resp := rest_resp.NewStatusCreated("success created category", result.Marshall(oauth.IsPublic(c.Request)))
	c.JSON(resp.Status(), resp)
}

func Get(c *gin.Context) {
	categoryID, err := controller_utils.GetIDInt(c.Param("category_id"), "category id")
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	result, getErr := services.CategoryService.GetCategory(categoryID)
	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}

	resp := rest_resp.NewStatusOK("success get category", result.Marshall(oauth.IsPublic(c.Request)))
	c.JSON(resp.Status(), resp)
}

func GetAll(c *gin.Context) {
	result, getErr := services.CategoryService.GetAllCategory()
	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}

	resp := rest_resp.NewStatusOK("success get category", result.Marshall(oauth.IsPublic(c.Request)))
	c.JSON(resp.Status(), resp)
}

func Update(c *gin.Context) {
	categoryID, err := controller_utils.GetIDInt(c.Param("category_id"), "category id")
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	var category categories.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	category.ID = categoryID
	result, updateErr := services.CategoryService.UpdateCategory(category)
	if updateErr != nil {
		c.JSON(updateErr.Status(), updateErr)
		return
	}

	resp := rest_resp.NewStatusOK("success updated category", result.Marshall(oauth.IsPublic(c.Request)))
	c.JSON(resp.Status(), resp)
}

func Delete(c *gin.Context) {
	categoryID, err := controller_utils.GetIDInt(c.Param("category_id"), "category id")
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	if err := services.CategoryService.DeleteCategory(categoryID); err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"message": "success deleted category", "status": http.StatusOK})
}
