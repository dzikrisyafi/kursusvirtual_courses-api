package enrolls

import (
	"net/http"
	"strconv"

	"github.com/dzikrisyafi/kursusvirtual_courses-api/src/domain/enrolls"
	"github.com/dzikrisyafi/kursusvirtual_courses-api/src/services"
	"github.com/dzikrisyafi/kursusvirtual_oauth-go/oauth"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_resp"
	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
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
		resp := rest_resp.NewStatusOK("success get enroll data", enroll.Marshall(false))
		c.JSON(resp.Status(), resp)
		return
	}

	resp := rest_resp.NewStatusOK("success get enroll data", enroll.Marshall(oauth.IsPublic(c.Request)))
	c.JSON(resp.Status(), resp)
}

func Create(c *gin.Context) {
	var request enrolls.Enroll
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	enroll, err := services.EnrollsService.CreateEnroll(request, c.Query("access_token"))
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, enroll)
}
