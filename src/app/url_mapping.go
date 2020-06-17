package app

import (
	"github.com/dzikrisyafi/kursusvirtual_courses-api/src/controllers/courses"
	"github.com/dzikrisyafi/kursusvirtual_courses-api/src/controllers/enrolls"
	"github.com/dzikrisyafi/kursusvirtual_middleware/middleware"
)

func mapUrls() {
	// course end point
	coursesGroup := router.Group("/courses")
	coursesGroup.Use(middleware.Auth())
	{
		coursesGroup.POST("/", courses.Create)
		coursesGroup.GET("/:course_id", courses.Get)
		coursesGroup.PUT("/:course_id", courses.Update)
		coursesGroup.PATCH("/:course_id", courses.Update)
		coursesGroup.DELETE("/:course_id", courses.Delete)
	}

	// internal course end point
	internalGroup := router.Group("/internal")
	internalGroup.Use(middleware.Auth())
	{
		internalGroup.GET("/courses/:course_name", courses.Search)
		internalGroup.GET("/enrolls/:user_id", enrolls.Get)
		internalGroup.POST("/enrolls", enrolls.Create)
		router.DELETE("/enrolls/:enroll_id", enrolls.Delete)
	}
}
