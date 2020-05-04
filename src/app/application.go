package app

import (
	"github.com/dzikrisyafi/kursusvirtual_courses-api/src/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()

	logger.Info("start the application...")
	router.Run(":8002")
}
