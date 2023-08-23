package main

import (
	"log/slog"
	"os"
	"time"
)

func main() {
	// slog.Debug("debug level")  // 1st level
	// slog.Info("info level")    // 2nd level
	// slog.Warn("warning level") // 3rd level
	// slog.Error("error level")  // 4th level

	// logHandlerTEXT := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
	// 	AddSource: true,
	// 	Level:     slog.LevelDebug,
	// })

	// logTEXT := slog.New(logHandlerTEXT)

	// logTEXT.Debug("What's the meaning of life?", slog.Int("answer", 42))

	logHandlerJSON := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			// Match the key that we want
			if a.Key == slog.TimeKey {
				a.Key = "date" // rename to date
				a.Value = slog.Int64Value(time.Now().Unix())
			}
			return a
		},
	}).WithAttrs([]slog.Attr{
		slog.Int("What's the meaning of life?", 42),
		slog.Group("votes", slog.Int("Picachu", 40), slog.Int("Mew", 24)),
	})

	logJSON := slog.New(logHandlerJSON)
	logJSON.Debug("Best Pokemon Rating")

	slog.SetDefault(logJSON)
	slog.Info("New Info")
}
