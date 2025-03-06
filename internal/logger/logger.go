package logger

import "log"

var reset = "\033[0m"
var red = "\033[31m"
var warning = "\033[33m"
var white = "\033[97m"

type Service interface {
	LogInfo(message string)
	LogWarning(message string)
	LogError(message string)
}

type Logger struct {
}

func NewLogger() Service {
	return &Logger{}
}

func (l *Logger) LogInfo(message string) {
	log.Println(white + message + reset)

}
func (l *Logger) LogWarning(message string) {
	log.Println(warning + message + reset)

}

func (l *Logger) LogError(message string) {
	log.Println(red + message + reset)
}
