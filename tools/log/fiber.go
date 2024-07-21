package log

import (
	"time"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/gofiber/fiber/v2"
)

func Fiber(l log.Logger) fiber.Handler {
  return func(c *fiber.Ctx) error {
    now := time.Now()
    err := c.Next()
    elapsed := time.Now().Sub(now)

    level.Info(l).Log(
      "msg", "request received",
      "method", c.Method(),
      "path", c.Path(),
      "body", len(c.BodyRaw()),
      "status", c.Response().StatusCode(),
      "duration", elapsed.Microseconds(),
      "userAgent", c.Get("User-Agent"),
    )

    return err
  }
}
