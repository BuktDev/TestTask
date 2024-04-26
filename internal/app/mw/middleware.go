package mw

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"strings"
)

const roleAdmin = "admin"

func RoleCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		val := c.Request().Header.Get("User-Role")

		if strings.EqualFold(val, roleAdmin) {
			log.Info().Msg("red button user detected")
		}

		err := next(c)
		if err != nil {
			return err
		}

		return nil
	}
}
