package courses

import (
	"net/http"
	"strconv"

	"github.com/dzikrisyafi/kursusvirtual_courses-api/src/domain/courses"
	"github.com/dzikrisyafi/kursusvirtual_courses-api/src/services"
	"github.com/dzikrisyafi/kursusvirtual_oauth-go/oauth"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
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

	c.JSON(http.StatusCreated, result.Marshall(oauth.IsPublic(c.Request)))
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

	c.JSON(http.StatusOK, course.Marshall(oauth.IsPublic(c.Request)))
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

	c.JSON(http.StatusOK, result.Marshall(oauth.IsPublic(c.Request)))
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

	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func Search(c *gin.Context) {
	courses, err := services.CoursesService.SearchCourse(c.Param("course_name"))
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, courses.Marshall(oauth.IsPublic(c.Request)))
}
