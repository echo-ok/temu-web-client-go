package log

import (
	"log/slog"
	"os"
	"testing"
)

func TestLogger(t *testing.T) {
	logger := NewSlogAdapter(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})))
	logger.Debugf("test\\ntest")
	println("test\\ntest")
}
