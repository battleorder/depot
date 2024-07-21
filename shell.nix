{ mkShell
, config

# Languages
, nodejs

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
    golangci-lint
    httpie
    just
    jq
    nodejs
    process-compose
    supabase-cli
  ];

  inputsFrom = [
    config.packages.units
  ];
}
