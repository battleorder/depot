package main

import (
	"github.com/battleorder/depot/tools/log"
	"github.com/battleorder/depot/units/internal/api"
	"github.com/battleorder/depot/units/internal/db"
	"github.com/go-kit/log/level"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func main() {
  logger := log.NewLogger()

  if err := db.Init(); err != nil {
    level.Error(logger).Log("msg", "failed to initialize supabase client", "err", err)
    panic(err)
  }

  app := fiber.New()
  app.Use(log.Fiber(logger))

  app.Use(cors.New(cors.Config{
    AllowOrigins: "http://localhost:5173",
    AllowCredentials: true,
    AllowHeaders: "Authorization",
  }))

  app.Use(limiter.New(limiter.Config{
    Storage: db.Storage,
  }))

  apiv1 := app.Group("/v1", api.Authenticatable)
  apiv1.Get("/units", api.ListUnits)
  apiv1.Post("/units", api.RequiresAuth, api.CreateUnit)

  level.Error(logger).Log("msg", "server crashed", "err", app.Listen(":4000"))
}
