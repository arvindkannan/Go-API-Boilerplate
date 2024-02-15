package utils

import (
	"os"
	"path/filepath"

	"log/slog"
)

func Logger() {
	level := slog.LevelInfo
	if debug := os.Getenv("SERVICE_DEBUG"); debug != "" {
		level = slog.LevelDebug
	}

	handler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level:     level,
		AddSource: true,
		ReplaceAttr: func(_ []string, attr slog.Attr) slog.Attr {
			if attr.Key == slog.SourceKey {
				source := attr.Value.Any().(*slog.Source)
				source.File = filepath.Base(source.File)
			}

			return attr
		},
	})

	slog.SetDefault(slog.New(handler))
}
