package helpers

import (
	"api-gateway/utils"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func AssertGrpcStatus(err error) error {
	if e, ok := status.FromError(err); ok {
		switch e.Code() {
		case codes.InvalidArgument:
			return echo.NewHTTPError(utils.ErrBadRequest.EchoFormatDetails(e.Message()))
		case codes.FailedPrecondition:
			return echo.NewHTTPError(utils.ErrBadRequest.EchoFormatDetails(e.Message()))
		case codes.NotFound:
			return echo.NewHTTPError(utils.ErrNotFound.EchoFormatDetails(e.Message()))
		case codes.Unauthenticated:
			return echo.NewHTTPError(utils.ErrUnauthorized.EchoFormatDetails(e.Message()))
		case codes.Internal:
			return echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(e.Message()))
		}
	}
	return echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(err.Error()))
}