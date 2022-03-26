package container

import (
	echostarter_wellknown "echo-starter/internal/wellknown"

	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/labstack/echo/v4"
)

func EnsureScopedContainer(root di.Container) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			subContainer, err := root.SubContainer()
			if err != nil {
				panic(err)
			}
			c.Set(echostarter_wellknown.SCOPED_CONTAINER_KEY, subContainer)
			return next(c)
		}
	}
}
