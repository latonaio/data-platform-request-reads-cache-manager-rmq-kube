package services

import (
	"github.com/astaxie/beego/context"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"golang.org/x/xerrors"
)

func VerifyToken(
	ctx *context.Context,
	l *logger.Logger,
	jwtToken string,
) {
	err := VerifyJwtToken(
		jwtToken,
		ctx,
	)

	l.Error(xerrors.Errorf("JWT Token is invalid: %w", err))
}
