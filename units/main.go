package main

import (
	"time"

	"github.com/battleorder/depot/tools/log"
	"github.com/battleorder/depot/units/internal/api"
	"github.com/battleorder/depot/units/internal/db"
	"github.com/go-kit/log/level"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/google/uuid"
)

func main() {
	logger := log.NewLogger()

	if err := db.Init(); err != nil {
		level.Error(logger).Log("msg", "failed to initialize supabase client", "err", err)
		panic(err)
	}

	app := fiber.New(fiber.Config{
    ErrorHandler: func(c *fiber.Ctx, err error) error {
      level.Error(logger).Log("msg", "request failed", "err", err)
      if err, ok := err.(*fiber.Error); ok {
        return c.Status(err.Code).SendString(err.Message)
      }
      return err
    },
  })
  
  app.Use(helmet.New())
  app.Use(etag.New())

  app.Use(requestid.New(requestid.Config{
    Header: "X-Request-ID",
    Generator: func() string {
      return uuid.NewString()
    },
  }))

	app.Use(log.Fiber(logger))

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowCredentials: true,
		AllowHeaders:     "Authorization",
	}))

	app.Use(limiter.New(limiter.Config{
		Max:        20,
		Expiration: 30 * time.Second,
		Storage:    db.Storage,
	}))

  app.Use(recover.New())

	apiv1 := app.Group("/v1", api.Authenticatable)
	apiv1.Get("/units", api.ListUnits)
	apiv1.Post("/units", api.RequiresAuth, api.CreateUnit)
  apiv1.Get("/units/:unitId", api.GetUnit)
  apiv1.Post("/units/:unitId/members", api.RequiresAuth, api.JoinUnit)

	level.Error(logger).Log("msg", "server crashed", "err", app.Listen(":4000"))
}
