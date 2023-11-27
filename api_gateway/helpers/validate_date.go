package helpers

import (
	"api-gateway/utils"
	"regexp"

	"github.com/labstack/echo/v4"
)

// Validate date with YYYY-MM-DD
func DateValidator(s string) error {
	r, _ := regexp.Compile(`^\d{4}\-(0[1-9]|1[012])\-(0[1-9]|[12][0-9]|3[01])$`)

	if !r.MatchString(s) {
		return echo.NewHTTPError(utils.ErrBadRequest.EchoFormatDetails("Invalid date format, date has to be YYYY-MM-DD"))
	}
	return nil
}