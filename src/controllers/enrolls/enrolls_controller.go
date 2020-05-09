package enrolls

import (
	"net/http"
	"strconv"

	"github.com/dzikrisyafi/kursusvirtual_courses-api/src/domain/enrolls"
	"github.com/dzikrisyafi/kursusvirtual_courses-api/src/services"
	"github.com/dzikrisyafi/kursusvirtual_oauth-go/oauth"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	if err := oauth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Status(), err)
		return
	}

	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		restErr := rest_errors.NewBadRequestError("user id should be a number")
		c.JSON(restErr.Status(), err)
		return
	}

	enroll, getErr := services.EnrollsService.GetCourseByUserID(userID)
	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}

	if oauth.GetCallerID(c.Request) == enroll.UserID {
		c.JSON(http.StatusOK, enroll.Marshall(false))
		return
	}

	c.JSON(http.StatusOK, enroll.Marshall(oauth.IsPublic(c.Request)))
}

func Create(c *gin.Context) {
	if err := oauth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Status(), err)
		return
	}

	if userID := oauth.GetCallerID(c.Request); userID == 0 {
		restErr := rest_errors.NewUnauthorizedError("invalid credentials")
		c.JSON(restErr.Status(), restErr)
	}

	var request enrolls.Enroll
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	enroll, err := services.EnrollsService.CreateEnroll(request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, enroll)
}
