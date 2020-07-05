package app

import (
	"github.com/dzikrisyafi/kursusvirtual_courses-api/src/controllers/cohort"
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
		coursesGroup.GET("/", courses.GetAll)
		coursesGroup.PUT("/:course_id", courses.Update)
		coursesGroup.PATCH("/:course_id", courses.Update)
		coursesGroup.DELETE("/:course_id", courses.Delete)
	}

	// internal course end point
	internalGroup := router.Group("/internal")
	internalGroup.Use(middleware.Auth())
	{
		internalGroup.GET("/courses/:course_name", courses.Search)

		internalGroup.POST("/enrolls", enrolls.Create)
		internalGroup.GET("/enrolls/:user_id", enrolls.Get)
		internalGroup.PUT("/enrolls/:enroll_id", enrolls.Update)
		internalGroup.DELETE("/enrolls/:enroll_id", enrolls.Delete)

		internalGroup.POST("/cohorts", cohort.Create)
		internalGroup.GET("/cohorts/:cohort_id", cohort.Get)
		internalGroup.GET("/cohorts", cohort.GetAll)
		internalGroup.PUT("/cohorts/:cohort_id", cohort.Update)
		internalGroup.DELETE("/cohorts/:cohort_id", cohort.Delete)
	}
}
