package mw

import (
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"
)

const (
	admin = "admin"
)

func RoleCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		userRole := ctx.Request().Header.Get("User-Role")

		if strings.EqualFold(userRole, admin) {
			fmt.Println("red button user detected")
		}

		err := next(ctx)
		if err != nil {
			return err
		}

		return nil
	}
}
