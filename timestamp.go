package logtail

import (
	"time"

	"github.com/rs/zerolog"
)

type LogtailTimestamp struct{}

func (LogtailTimestamp) Run(e *zerolog.Event, level zerolog.Level, message string) {
	e.Time("dt", time.Now().UTC())
}
