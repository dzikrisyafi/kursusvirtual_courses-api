package courses

import (
	"net/http"

	"github.com/dzikrisyafi/kursusvirtual_courses-api/src/domain/courses"
	"github.com/dzikrisyafi/kursusvirtual_courses-api/src/services"
	"github.com/dzikrisyafi/kursusvirtual_oauth-go/oauth"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/controller_utils"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_resp"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var course courses.Course
	if err := c.ShouldBindJSON(&course); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	result, err := services.CoursesService.CreateCourse(course)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	resp := rest_resp.NewStatusCreated("success created course", result.Marshall(oauth.IsPublic(c.Request)))
	c.JSON(resp.Status(), resp)
}

func Get(c *gin.Context) {
	courseID, err := controller_utils.GetIDInt(c.Param("course_id"), "course id")
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	result, getErr := services.CoursesService.GetCourse(courseID)
	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}

	resp := rest_resp.NewStatusOK("success get course", result.Marshall(oauth.IsPublic(c.Request)))
	c.JSON(resp.Status(), resp)
}

func GetAll(c *gin.Context) {
	result, getErr := services.CoursesService.GetAllCourse()
	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
	}

	resp := rest_resp.NewStatusOK("success get course", result.Marshall(oauth.IsPublic(c.Request)))
	c.JSON(resp.Status(), resp)
}

func Update(c *gin.Context) {
	courseID, err := controller_utils.GetIDInt(c.Param("course_id"), "course id")
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	var course courses.Course
	if err := c.ShouldBindJSON(&course); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	course.ID = courseID
	isPartial := c.Request.Method == http.MethodPatch
	result, updateErr := services.CoursesService.UpdateCourse(isPartial, course)
	if err != nil {
		c.JSON(updateErr.Status(), updateErr)
		return
	}

	resp := rest_resp.NewStatusOK("success updating course", result.Marshall(oauth.IsPublic(c.Request)))
	c.JSON(resp.Status(), resp)
}

func Delete(c *gin.Context) {
	courseID, err := controller_utils.GetIDInt(c.Param("course_id"), "course id")
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	if err := services.CoursesService.DeleteCourse(courseID, c.Query("access_token")); err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"message": "success deleted course", "status": http.StatusOK})
}

func Search(c *gin.Context) {
	courses, err := services.CoursesService.SearchCourse(c.Param("course_name"))
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	resp := rest_resp.NewStatusOK("success get course", courses.Marshall(oauth.IsPublic(c.Request)))
	c.JSON(resp.Status(), resp)
}
