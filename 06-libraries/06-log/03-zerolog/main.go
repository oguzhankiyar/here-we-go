package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

func main() {
	Sample("Simple", Simple)
	Sample("Level", Level)
	Sample("Params", Params)
	Sample("Send", Send)
	Sample("Format", Format)
	Sample("FieldName", FieldName)
}

func Simple() {
	logger := zerolog.New(os.Stdout)
	logger.Debug().Msg("this is simple log")
}

func Level() {
	logger := zerolog.New(os.Stdout).Level(zerolog.WarnLevel)
	logger.Info().Msg("this is info log")
	logger.Warn().Msg("this is warn log")
	logger.Debug().Msg("this is debug log")
}

func Params() {
	logger := zerolog.New(os.Stdout)
	logger.
		Debug().
		Str("user", "gopher").
		Int("age", 10).
		Msg("this is log with params")
}

func Send() {
	logger := zerolog.New(os.Stdout)
	logger.
		Debug().
		Str("name", "gopher").
		Send()
}

func Format() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	logger := zerolog.New(os.Stdout)

	err := errors.New("an error")
	logger.Error().Stack().Err(err).Msg("")
}

func FieldName() {
	zerolog.TimestampFieldName = "t"
	zerolog.LevelFieldName = "l"
	zerolog.MessageFieldName = "m"

	logger := zerolog.New(os.Stdout)
	logger.Info().Msg("this is an info")
}

func Sample(name string, fn func()) {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}