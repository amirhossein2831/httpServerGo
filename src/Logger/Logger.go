package Logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"sync"
)

var (
	loggerInstance Logger
	once           sync.Once
)

type Logger interface {
	GetLogger() *zap.Logger
	SetLogger(logger *zap.Logger)
}

type Log struct {
	logger *zap.Logger
}

func GetInstance() Logger {
	once.Do(func() {
		core := zapcore.NewTee(zapcore.NewCore(encoderFile(), logFile(), zapcore.DebugLevel))
		loggerInstance = &Log{
			logger: zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel)),
		}
	})
	return loggerInstance
}

func encoderFile() zapcore.Encoder {
	conf := zap.NewProductionEncoderConfig()
	conf.EncodeTime = zapcore.ISO8601TimeEncoder
	return zapcore.NewJSONEncoder(conf)
}

func logFile() zapcore.WriteSyncer {
	file, _ := os.OpenFile("logs/log.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	return zapcore.AddSync(file)
}

func (l *Log) GetLogger() *zap.Logger {
	return l.logger
}

func (l *Log) SetLogger(logger *zap.Logger) {
	l.logger = logger
}
