package main

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	Sample("Development", Development)
	Sample("Production", Production)
	Sample("Sugar", Sugar)
	Sample("With", With)
	Sample("Simple", Simple)
	Sample("Level", Level)
}

func Development() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	logger.Info("this is development logger")
}

func Production() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	logger.Info("this is production logger")
}

func Sugar() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	sugar := logger.Sugar()

	sugar.Info("this is sugar logger")
}

func With() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	with := logger.With(
		zap.String("version", "1.0.0"),
		zap.String("app", "zap-sample"),
		zap.String("user", "gopher"))

	with.Info("this is with logger")
}

func Simple() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	sugar := logger.Sugar()

	sugar.Info("this is simple log")
	sugar.Infow("this is simple log with args",
		"user", "gopher")
}

func Level() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	logger.Core().Enabled(zapcore.WarnLevel)

	logger.Debug("this is debug log")
	logger.Info("this is info log")
}

func Sample(name string, fn func()) {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}