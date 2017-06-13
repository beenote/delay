package delay

import (
	"time"

	"github.com/labstack/echo"
)

// DelayHeader can have value like 300ms 2.5s and 0.5m
const DelayHeader = "X-Add-Delay"

// New add delays latency to your endpoint
func New() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			v := c.Request().Header.Get(DelayHeader)
			if v != "" {
				d, err := time.ParseDuration(v)
				if err == nil {
					time.Sleep(d)
				}
			}
			return next(c)
		}
	}
}
