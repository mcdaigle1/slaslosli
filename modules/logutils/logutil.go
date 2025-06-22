package logutils

import (
	"log"
	"log/slog"
	"os"

	"github.com/mcdaigle1/slaslosli/config"
)

var logLevels = map[string]slog.Level{
    "DEBUG": slog.LevelDebug,
    "INFO":  slog.LevelInfo,
    "WARN":  slog.LevelWarn,
    "ERROR": slog.LevelError,
}

var (
    // Exported so other packages can update the log level
    LogLevel = new(slog.LevelVar)
    logger   *slog.Logger
)

// InitLogger initializes the global logger with dynamic log level.
func InitLogger() error {
    configLogLevel := config.Global.LogLevel
    if sLogLevel, ok := logLevels[configLogLevel]; ok {
        log.Printf("Setting log level to: %s", configLogLevel)
        LogLevel.Set(sLogLevel)
    } else {
        log.Printf("Could not find log level: %s, using INFO log level", configLogLevel)
        LogLevel.Set(slog.LevelInfo)
    }

    handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
        Level: LogLevel,
    })

    logger = slog.New(handler)
    slog.SetDefault(logger)

    return nil
}
