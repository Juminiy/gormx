package logger

import (
	"gorm.io/gorm/logger"
	"time"
)

// gorm no color logger interface
// gorm.Logger

func Plain(printfFn func(string, ...any)) logger.Interface {
	return logger.New(
		printfWriter{f: printfFn},
		logger.Config{
			SlowThreshold:             200 * time.Millisecond,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: false,
			Colorful:                  false,
		})
}

type printfWriter struct {
	f func(string, ...any)
}

func (w printfWriter) Printf(s string, v ...any) {
	w.f(s, v...)
}
