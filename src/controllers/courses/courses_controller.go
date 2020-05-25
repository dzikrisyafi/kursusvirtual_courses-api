package courses

import (
	"net/http"
	"strconv"

	"github.com/dzikrisyafi/kursusvirtual_courses-api/src/domain/courses"
	"github.com/dzikrisyafi/kursusvirtual_courses-api/src/services"
	"github.com/dzikrisyafi/kursusvirtual_oauth-go/oauth"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_resp"
	"github.com/gin-gonic/gin"
)

func ParseId(courseIdParam string) (int64, rest_errors.RestErr) {
	courseID, err := strconv.ParseInt(courseIdParam, 10, 64)
	if err != nil {
		return 0, rest_errors.NewBadRequestError("course id should be a number")
	}

	return courseID, nil
}

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

	resp := rest_resp.NewStatusCreated("success creating course", result.Marshall(oauth.IsPublic(c.Request)))
	c.JSON(resp.Status(), resp)
}

func Get(c *gin.Context) {
	courseID, err := ParseId(c.Param("course_id"))
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	course, getErr := services.CoursesService.GetCourse(courseID)
	if err != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}

	resp := rest_resp.NewStatusOK("success creating course", course.Marshall(oauth.IsPublic(c.Request)))
	c.JSON(resp.Status(), resp)
}

func Update(c *gin.Context) {
	courseID, err := ParseId(c.Param("course_id"))
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
	courseID, err := ParseId(c.Param("course_id"))
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	if err := services.CoursesService.DeleteCourse(courseID); err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"message": "success deleting course", "status": http.StatusOK})
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
