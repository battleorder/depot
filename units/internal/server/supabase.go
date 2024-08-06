package server

import (
	"context"
	"log/slog"
	"net/http"
	"strings"

	"github.com/battleorder/depot/units/internal/db"
	"github.com/supabase-community/gotrue-go"
	"github.com/supabase-community/gotrue-go/types"
	"github.com/supabase-community/supabase-go"
)

type RequestKey int

const (
	supabaseAuthClientKey RequestKey = iota
	supabaseClientKey     RequestKey = iota + 1
	authTokenKey          RequestKey = iota + 2
	authUserKey           RequestKey = iota + 3
)

func newAuthenticatedMiddleware(spb *supabase.Client, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		authzHeader := r.Header.Get("Authorization")
		if strings.HasPrefix(authzHeader, "Bearer ") {
			// Pull token from headers
			token := strings.Split(authzHeader, " ")[1]
			ctx = context.WithValue(ctx, authTokenKey, token)

			// Create a Supabase client with that token
			client := spb.Auth.WithToken(token)
			ctx = context.WithValue(ctx, supabaseAuthClientKey, client)

			// Get the user with the new Supabase client
			user, err := client.GetUser()
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				if !strings.Contains(err.Error(), "token is expired by") {
          slog.Error("failed to retrieve user", "err", err)
				}
				return
			}
			ctx = context.WithValue(ctx, authUserKey, user) //nolint:ineffassign
		}

		// Pass to the next handler in the stack. If the context hasn't changed this
		// largely does nothing.
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

func RequiresAuth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if user, ok := GetUser(r); !ok || user == nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		h.ServeHTTP(w, r)
	})
}

func GetSupabaseAuthClient(r *http.Request) (gotrue.Client, bool) {
	client, ok := r.Context().Value(supabaseAuthClientKey).(gotrue.Client)
	return client, ok
}

func GetSupabaseClient(r *http.Request) (*supabase.Client, error) {
	// No token = use an anonymous client
	client, err := db.NewClient()
	if err != nil {
		return nil, err
	}

	token, ok := GetToken(r)
	if !ok {
		return client, err
	}

	client.UpdateAuthSession(types.Session{
		AccessToken: token,
	})

	return client, err
}

func GetUser(r *http.Request) (*types.UserResponse, bool) {
	user, ok := r.Context().Value(authUserKey).(*types.UserResponse)
	return user, ok
}

func GetToken(r *http.Request) (string, bool) {
	token, ok := r.Context().Value(authTokenKey).(string)
	return token, ok
}
