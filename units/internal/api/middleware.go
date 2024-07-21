package api

import (
	"strings"

	"github.com/battleorder/depot/units/internal/db"
	"github.com/gofiber/fiber/v2"
	"github.com/supabase-community/gotrue-go"
	"github.com/supabase-community/gotrue-go/types"
)

func Authenticatable(c *fiber.Ctx) error {
  authzHeader := c.Get("Authorization")
  if strings.HasPrefix(authzHeader, "Bearer ") {
    // Authn is being attempted
    token := strings.Split(authzHeader, " ")[1] // get token
    withToken := db.Client.Auth.WithToken(token)
    c.Locals("sb-auth-client", withToken)
    userData, err := withToken.GetUser()
    if err != nil {
      return err
    }
    c.Locals("sb-auth-user", userData)
  }

  return c.Next()
}

func RequiresAuth(c *fiber.Ctx) error {
  if user, ok := GetAuthUser(c); !ok || user == nil {
    return fiber.ErrUnauthorized
  }
  return c.Next()
}

func GetAuthClient(c *fiber.Ctx) (gotrue.Client, bool) {
  client, ok := c.Locals("sb-auth-client").(gotrue.Client)
  return client, ok
}

func GetAuthUser(c *fiber.Ctx) (*types.UserResponse, bool) {
  user, ok := c.Locals("sb-auth-user").(*types.UserResponse)
  return user, ok
}
