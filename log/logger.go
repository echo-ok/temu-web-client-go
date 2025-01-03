package log

import (
	"log/slog"
)

type SlogAdapter struct {
	logger *slog.Logger
}

func NewSlogAdapter(logger *slog.Logger) *SlogAdapter {
	return &SlogAdapter{logger: logger}
}

func (s *SlogAdapter) Errorf(format string, v ...interface{}) {
	s.logger.Error(format, v...)
}

func (s *SlogAdapter) Warnf(format string, v ...interface{}) {
	s.logger.Warn(format, v...)
}

func (s *SlogAdapter) Debugf(format string, v ...interface{}) {
	s.logger.Debug(format, v...)
}
