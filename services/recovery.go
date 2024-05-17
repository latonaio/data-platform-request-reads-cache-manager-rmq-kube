package services

import (
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"runtime/debug"
)

func Recover(
	customLogger *logger.Logger,
	controller *beego.Controller,
) {
	if err := recover(); err != nil {
		customLogger.Error(err)
		debug.PrintStack()
		if controller != nil {
			status := 500
			HandleError(
				controller,
				err,
				&status,
			)
		}
	}
}
