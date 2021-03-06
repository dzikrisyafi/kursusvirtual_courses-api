package enrolls

import (
	"net/http"

	"github.com/dzikrisyafi/kursusvirtual_courses-api/src/domain/enrolls"
	"github.com/dzikrisyafi/kursusvirtual_courses-api/src/services"
	"github.com/dzikrisyafi/kursusvirtual_oauth-go/oauth"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/controller_utils"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_resp"
	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	userID, err := controller_utils.GetIDInt(c.Param("user_id"), "user id")
	if err != nil {
		restErr := rest_errors.NewBadRequestError("user id should be a number")
		c.JSON(restErr.Status(), err)
		return
	}

	result, getErr := services.EnrollsService.GetCourseByUserID(userID)
	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}

	if oauth.GetCallerID(c.Request) == result.UserID {
		resp := rest_resp.NewStatusOK("success get enroll data", result.Marshall(false))
		c.JSON(resp.Status(), resp)
		return
	}

	resp := rest_resp.NewStatusOK("success get enroll data", result.Marshall(oauth.IsPublic(c.Request)))
	c.JSON(resp.Status(), resp)
}

func Create(c *gin.Context) {
	var request enrolls.Enroll
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	result, err := services.EnrollsService.CreateEnroll(request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func Update(c *gin.Context) {
	enrollID, err := controller_utils.GetIDInt(c.Param("enroll_id"), "enroll id")
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	var request enrolls.Enroll
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	request.ID = enrollID
	result, err := services.EnrollsService.UpdateEnroll(request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func Delete(c *gin.Context) {
	enrollID, idErr := controller_utils.GetIDInt(c.Param("enroll_id"), "enroll id")
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	if err := services.EnrollsService.DeleteEnroll(enrollID); err != nil {
		c.JSON(err.Status(), err)
		return
	}
}
