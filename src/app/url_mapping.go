package app

import (
	"github.com/dzikrisyafi/kursusvirtual_courses-api/src/controllers/courses"
	"github.com/dzikrisyafi/kursusvirtual_courses-api/src/controllers/enrolls"
)

func mapUrls() {
	router.POST("/courses", courses.Create)
	router.GET("/courses/:course_id", courses.Get)
	router.PUT("/courses/:course_id", courses.Update)
	router.PATCH("/courses/:course_id", courses.Update)
	router.DELETE("/courses/:course_id", courses.Delete)

	router.GET("/internal/courses/:course_name", courses.Search)
	router.GET("/internal/enrolls/:user_id", enrolls.Get)
	router.POST("/internal/enrolls", enrolls.Create)
}
