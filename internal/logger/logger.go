package logger

import (
	"context"
	"log/slog"
)

type Service interface {
	LogInfo(message string)
	LogSuccess(message string)
	LogWarning(message string)
	LogError(message string)
}

type Colors struct {
	reset   string
	red     string
	warning string
	white   string
	green   string
}

type Logger struct {
	colors Colors
	logger *slog.Logger
}

func NewLogger() Service {
	return &Logger{
		logger: slog.Default(),
		colors: Colors{
			reset:   "\u001B[0m",
			red:     "\u001B[31m",
			warning: "\u001B[33m",
			white:   "\u001B[97m",
			green:   "\u001B[32m",
		},
	}
}

func (l *Logger) LogSuccess(message string) {
	l.logger.Log(context.Background(), slog.LevelInfo, l.colors.green+message+l.colors.reset)
}

func (l *Logger) LogInfo(message string) {
	l.logger.Log(context.Background(), slog.LevelInfo, l.colors.white+message+l.colors.reset)
}

func (l *Logger) LogWarning(message string) {
	l.logger.Log(context.Background(), slog.LevelWarn, l.colors.warning+message+l.colors.reset)
}

func (l *Logger) LogError(message string) {
	l.logger.Log(context.Background(), slog.LevelError, l.colors.red+message+l.colors.reset)
}
