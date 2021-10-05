package main

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
)

func main() {
	Sample("Default", Default)
	Sample("JSONFormatter", JSONFormatter)
	Sample("TextFormatter", TextFormatter)
	Sample("SetOutput", SetOutput)
	Sample("SetLevel", SetLevel)
	Sample("WithFields", WithFields)
	Sample("ContextLogger", ContextLogger)
}

func Default() {
	logger :=  logrus.New()

	logger.Println("this is a simple log")
}

func JSONFormatter() {
	logger :=  logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.Println("this is a json log")
}

func TextFormatter() {
	logger :=  logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{})
	logger.Println("this is a text log")
}

func SetOutput() {
	var buff bytes.Buffer

	logger :=  logrus.New()
	logger.SetOutput(&buff)
	logger.Println("this is a log")

	fmt.Println(buff.String())
}

func SetLevel() {
	logger :=  logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.WarnLevel)
	logger.Println("this is a println log")
	logger.Debug("this is a debug log")
	logger.Warn("this is a warn log")
	logger.Info("this is a info log")
}

func WithFields() {
	logger :=  logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.WithFields(logrus.Fields{
		"name": "gopher",
		"age":   10,
	}).Info("gopher is here")
}

func ContextLogger() {
	logger :=  logrus.New()

	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.Info("hello")

	contextLogger := logger.WithFields(logrus.Fields{
		"version": "1.0.0",
		"app": "logrus-sample",
		"user": "gopher",
	})

	contextLogger.Info("the app is started")
	contextLogger.Warn("config is not found, using default values")
}

func Sample(name string, fn func()) {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}