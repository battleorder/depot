package db

import (
	"os"

	"github.com/supabase-community/supabase-go"
)

var (
	Client    *supabase.Client
	ApiUrl    = os.Getenv("SUPABASE_API_URL")
	AnonToken = os.Getenv("SUPABASE_ANON_KEY")
)

func Init() error {
  c, err := NewClient()
	if err != nil {
		return err
	}
	Client = c

	initRedis()

	return nil
}

func NewClient() (*supabase.Client, error) {
  return supabase.NewClient(ApiUrl, AnonToken, &supabase.ClientOptions{
		Schema: "units",
	})
}
