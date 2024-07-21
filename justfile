lint-go:
  golangci-lint run ./...

lint-db:
  supabase db lint

lint: lint-go lint-db
  @echo ""
  @echo "✨ Linting completed! ✨"

get-api-tokens user='user1':
  http POST \
    127.0.0.1:54321/auth/v1/token?grant_type=password \
    email={{user}}@battleorder.me \
    password=user123! \
    | jq -r '{ accessToken: .access_token, refreshToken: .refresh_token }'
