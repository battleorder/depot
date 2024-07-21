package log

import (
	"os"

	gokit_log "github.com/go-kit/log"
)

func NewLogger() gokit_log.Logger {
  l := gokit_log.NewLogfmtLogger(os.Stderr)
  l = gokit_log.WithPrefix(l, "ts", gokit_log.DefaultTimestamp)
  return l
}
