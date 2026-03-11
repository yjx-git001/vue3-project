package logger

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjackv2 "gopkg.in/natefinch/lumberjack.v2"
)

const (
	Console    = "console"
	File       = "file"
	TrafficKey = "X-Request-Id"
)

var (
	Level = zap.DebugLevel
	//Level = zap.InfoLevel
	Target = Console
)

var (
	Logger *zap.Logger
	Sugar  *zap.SugaredLogger
	core   zapcore.Core
)

func NewEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		// Keys can be anything except the empty string.
		TimeKey:        "T",
		LevelKey:       "L",
		NameKey:        "N",
		CallerKey:      "C",
		MessageKey:     "M",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

func init() {
	w := zapcore.AddSync(&lumberjackv2.Logger{
		Filename:   "log/dt_server.log",
		MaxSize:    1024, // megabytes
		MaxBackups: 10,
		MaxAge:     7, // days
	})

	var writeSyncer zapcore.WriteSyncer
	if Target == Console {
		writeSyncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout))
	}
	if Target == File {
		writeSyncer = zapcore.NewMultiWriteSyncer(w)
	}

	core = zapcore.NewCore(
		zapcore.NewConsoleEncoder(NewEncoderConfig()),
		writeSyncer,
		Level,
	)
	Logger = zap.New(core, zap.AddCaller())
	Sugar = Logger.Sugar()
}

func NewRequestIdLogger(requestId string) *zap.SugaredLogger {
	return zap.New(core, zap.AddCaller()).With(zap.String(TrafficKey, requestId)).Sugar()
}
