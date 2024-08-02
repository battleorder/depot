package log

import (
	"io"

	gokit_log "github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

func NewLogger(w io.Writer) gokit_log.Logger {
  lgr := gokit_log.NewLogfmtLogger(w)
  lgr = gokit_log.WithPrefix(lgr, "ts", gokit_log.DefaultTimestamp)
  lgr = level.NewFilter(lgr, level.AllowAll())
  lgr = gokit_log.WithPrefix(lgr, "module", gokit_log.DefaultCaller)
  return lgr
}
