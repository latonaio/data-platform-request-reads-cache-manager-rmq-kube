package services

import (
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func Recover(
	customLogger *logger.Logger,
) {
	if err := recover(); err != nil {
		customLogger.Error(err)
	}
}
