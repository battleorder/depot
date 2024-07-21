package main

import (
	"log"

	"github.com/battleorder/depot/units/internal/api"
	"github.com/battleorder/depot/units/internal/db"
	"github.com/gofiber/fiber/v2"
)

func main() {
  if err := db.Init(); err != nil {
    panic(err)
  }

  app := fiber.New()
  apiv1 := app.Group("/v1", api.Authenticatable)
  apiv1.Get("/units", api.ListUnits)
  apiv1.Post("/units", api.RequiresAuth, api.CreateUnit)

  log.Fatal(app.Listen(":4000"))
}
