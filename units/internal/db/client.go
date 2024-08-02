package db

import (
	"github.com/supabase-community/supabase-go"
)

var (
	ApiUrl    string
	AnonToken string
)

func NewSupabase(
  apiUrl string,
  anonToken string,
) (*supabase.Client, error) {
	ApiUrl = apiUrl
	AnonToken = anonToken

	c, err := NewClient()
	if err != nil {
		return nil, err
	}

	initRedis()

	return c, nil
}

func NewClient() (*supabase.Client, error) {
	return supabase.NewClient(ApiUrl, AnonToken, &supabase.ClientOptions{
		Schema: "units",
	})
}
