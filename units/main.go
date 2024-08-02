package main

import (
  "flag"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/battleorder/depot/tools/log"
	"github.com/battleorder/depot/units/internal/db"
	"github.com/battleorder/depot/units/internal/server"
	"github.com/go-kit/log/level"
)

func run(ctx context.Context, w io.Writer, getenv func(string) string, args []string) error {
  ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
  defer cancel()

  fs := flag.NewFlagSet("server", flag.ContinueOnError)

  addr := fs.String("server-addr", "0.0.0.0:4000", "the address to serve HTTP requests on")
  supabaseApiUrl := fs.String("supabase-url", getenv("SUPABASE_API_URL"), "the API URL of the supabase instance (uses $SUPABASE_API_URL if not provided)")
  supabaseAnonKey := fs.String("supabase-anon-key", getenv("SUPABASE_ANON_KEY"), "the API key of the anon role in the supabase instance (uses $SUPABASE_ANON_KEY if not provided)")

  if err := fs.Parse(args[1:]); err != nil {
    return err
  }

  lgr := log.NewLogger(w)

  spb, err := db.NewSupabase(*supabaseApiUrl, *supabaseAnonKey)
  if err != nil {
    level.Error(lgr).Log("msg", "error setting up supabase client", "err", err)
    return err
  }

  srv := server.NewServer(lgr, spb)

  httpServer := &http.Server{
    Addr: *addr,
    Handler: srv,
  }

  go func () {
    level.Info(lgr).Log("msg", "listening http server", "addr", httpServer.Addr)
    if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
      level.Error(lgr).Log("msg", "error listening and serving", "err", err)
    }
  }()

  var wg sync.WaitGroup
  wg.Add(1)
  go func() {
    defer wg.Done()
    <-ctx.Done()
    //nolint:ineffassign,staticcheck
    shutdownCtx := context.Background()
    shutdownCtx, cancel := context.WithTimeout(ctx, 10 * time.Second)
    defer cancel()
    if err := httpServer.Shutdown(shutdownCtx); err != nil {
      level.Error(lgr).Log("msg", "error shutting down http server", "err", err)
    }
  }()
  wg.Wait()
  return nil
}

func main() {
  ctx := context.Background()
  if err := run(ctx, os.Stdout, os.Getenv, os.Args); err != nil {
    fmt.Fprintf(os.Stderr, "%s\n", err)
    os.Exit(1)
  }
}
