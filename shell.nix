{ mkShell

# Languages
, nodejs
, go

# Scripting
, just
, process-compose
, jq

# Dev Env Tools
, supabase-cli
, golangci-lint
, httpie
}:
mkShell {
  buildInputs = [
    go
    golangci-lint
    httpie
    just
    jq
    nodejs
    process-compose
    supabase-cli
  ];
}
