package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"cron-sample/internal/infrastructure/config/models"
	"cron-sample/internal/infrastructure/logger/interfaces"
)

type AppLogger struct {
	logName  string
	logLevel string
	devMode  bool
	encoding string
	logger   *zap.Logger
}

func NewAppLogger(config models.LoggerConfig) interfaces.Logger {
	return &AppLogger{
		logName:  config.LogName,
		logLevel: config.LogLevel,
		devMode:  config.DevMode,
		encoding: config.Encoder,
	}
}

var loggerLevelMap = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
	"panic": zapcore.PanicLevel,
	"fatal": zapcore.FatalLevel,
}

func (l *AppLogger) Init() error {
	logLevel := zapcore.DebugLevel
	if level, exist := loggerLevelMap[l.logLevel]; exist {
		logLevel = level
	}

	logWriter := zapcore.AddSync(os.Stdout)

	var encoderCfg zapcore.EncoderConfig
	if l.devMode {
		encoderCfg = zap.NewDevelopmentEncoderConfig()
	} else {
		encoderCfg = zap.NewProductionEncoderConfig()
	}

	var encoder zapcore.Encoder
	encoderCfg.NameKey = "service"
	encoderCfg.TimeKey = "time"
	encoderCfg.LevelKey = "level"
	encoderCfg.FunctionKey = "function"
	encoderCfg.CallerKey = "caller"
	encoderCfg.MessageKey = "message"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderCfg.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderCfg.EncodeCaller = zapcore.ShortCallerEncoder
	encoderCfg.EncodeName = zapcore.FullNameEncoder
	encoderCfg.EncodeDuration = zapcore.StringDurationEncoder

	if l.encoding == "console" {
		encoder = zapcore.NewConsoleEncoder(encoderCfg)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderCfg)
	}

	core := zapcore.NewCore(encoder, logWriter, zap.NewAtomicLevelAt(logLevel))
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	l.logger = logger.Named(l.logName)

	go l.logger.Sync()

	return nil
}

func (l *AppLogger) Debug(msg string, args ...map[string]interface{}) {
	l.logger.Debug(msg, *mapFields(args...)...)
}

func (l *AppLogger) Info(msg string, args ...map[string]interface{}) {
	l.logger.Info(msg, *mapFields(args...)...)
}

func (l *AppLogger) Warn(msg string, args ...map[string]interface{}) {
	l.logger.Warn(msg, *mapFields(args...)...)
}

func (l *AppLogger) Error(msg string, err error, args ...map[string]interface{}) {
	fields := *mapFields(args...)
	fields = append(fields, zap.Error(err))

	l.logger.Error(msg, fields...)
}

func (l *AppLogger) Fatal(msg string, err error, args ...map[string]interface{}) {
	fields := *mapFields(args...)
	fields = append(fields, zap.Error(err))

	l.logger.Fatal(msg, fields...)
}

func mapFields(parameters ...map[string]interface{}) *[]zap.Field {
	var fields []zap.Field

	for _, parameterItem := range parameters {
		for key, value := range parameterItem {
			fields = append(fields, zap.Any(key, value))
		}
	}

	return &fields
}
