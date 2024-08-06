package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/battleorder/depot/units/internal/db"
	"github.com/battleorder/depot/units/internal/server"
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

	lgr1 := slog.New(
    slog.NewTextHandler(w, &slog.HandlerOptions{
      Level: slog.LevelDebug,
      AddSource: true,
    }),
  )
	slog.SetDefault(lgr1)

	spb, err := db.NewSupabase(*supabaseApiUrl, *supabaseAnonKey)
	if err != nil {
    slog.Error("error setting up supabase client", "err", err)
		return err
	}

	srv := server.NewServer(spb)

	httpServer := &http.Server{
		Addr:    *addr,
		Handler: srv,
	}

	go func() {
    slog.Info("listening http server", "addr", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
      slog.Error("error listening and serving", "err", err)
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		//nolint:ineffassign,staticcheck
		shutdownCtx := context.Background()
		shutdownCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
      slog.Error("error shutting down http server", "err", err)
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
