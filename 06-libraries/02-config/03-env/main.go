package main

import (
	"fmt"
	"time"

	"github.com/caarlos0/env/v6"
)

func main() {
	Sample("Simple", Simple)
	Sample("Options", Options)
	Sample("File", File)
}

func Simple() {
	type config struct {
		Home         string        `env:"HOME"`
		Port         int           `env:"PORT" envDefault:"3000"`
		Password     string        `env:"PASSWORD,unset"`
		IsProduction bool          `env:"PRODUCTION"`
		Hosts        []string      `env:"HOSTS" envSeparator:":"`
		Duration     time.Duration `env:"DURATION"`
		TempFolder   string        `env:"TEMP_FOLDER" envDefault:"${HOME}/tmp" envExpand:"true"`
	}

	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	fmt.Printf("%+v\n", cfg)
}

func Options() {
	type Config struct {
		Password string `env:"PASSWORD"`
	}

	cfg := &Config{}
	opts := env.Options{Environment: map[string]string{
		"PASSWORD": "MY_PASSWORD",
	}}

	if err := env.Parse(cfg, opts); err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("%+v\n", cfg.Password)
}

func File() {
	type Config struct {
		Password     string   `env:"PASSWORD,file" envDefault:"./06-libraries/02-config/03-ENV/config/password.txt"`
	}

	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	fmt.Printf("%+v\n", cfg)
}

func Sample(name string, fn func())  {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}