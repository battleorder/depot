package db

import (
	"github.com/supabase-community/supabase-go"
)

var (
	Client    *supabase.Client
	ApiUrl    = "http://127.0.0.1:54321"
	AnonToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZS1kZW1vIiwicm9sZSI6InNlcnZpY2Vfcm9sZSIsImV4cCI6MTk4MzgxMjk5Nn0.EGIM96RAZx35lJzdJsyH-qQwv8Hdp7fsn3W0YpN81IU"
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
