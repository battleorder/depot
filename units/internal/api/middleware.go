package api

import (
	"strings"

	"github.com/battleorder/depot/units/internal/db"
	"github.com/gofiber/fiber/v2"
	"github.com/supabase-community/gotrue-go"
	"github.com/supabase-community/gotrue-go/types"
	"github.com/supabase-community/supabase-go"
)

func Authenticatable(c *fiber.Ctx) error {
	authzHeader := c.Get("Authorization")
	if strings.HasPrefix(authzHeader, "Bearer ") {
		// Authn is being attempted
		token := strings.Split(authzHeader, " ")[1] // get token
		c.Locals("sb-auth-token", token)

		// Authn is being pushed to Supabase
		withToken := db.Client.Auth.WithToken(token)
		c.Locals("sb-auth-client", withToken)

		// Authn should now be successul
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

func GetUserSupabase(c *fiber.Ctx) (*supabase.Client, error) {
	accessToken, ok := c.Locals("sb-auth-token").(string)
	if !ok {
		return nil, nil
	}

	client, err := db.NewClient()
	if err != nil {
		return nil, err
	}

  client.UpdateAuthSession(types.Session{
    AccessToken: accessToken,
  })

	return client, nil
}

func GetAuthUser(c *fiber.Ctx) (*types.UserResponse, bool) {
	user, ok := c.Locals("sb-auth-user").(*types.UserResponse)
	return user, ok
}
