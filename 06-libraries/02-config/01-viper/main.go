package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func main() {
	Sample("Simple", Simple)
	Sample("Unmarshal", Unmarshal)
	Sample("Env", Env)
	Sample("New", New)
	Sample("Write", Write)
	Sample("Keys", Keys)
	Sample("Settings", Settings)
	Sample("Watch", Watch)
}

func Simple() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./06-libraries/02-config/01-viper/config")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("appName:", viper.Get("appName"))

	viper.Reset()
}

func Unmarshal() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./06-libraries/02-config/01-viper/config")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("error:", err)
		return
	}

	type Config struct {
		AppName	string
	}

	var config Config

	err := viper.Unmarshal(&config)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("appName:", config.AppName)

	viper.Reset()
}

func Env() {
	os.Setenv("GO_APP_NAME", "Here We Go")

	viper.AutomaticEnv()

	fmt.Println("appName:", viper.Get("GO_APP_NAME"))

	os.Remove("GO_APP_NAME")
	viper.Reset()
}

func New() {
	v := viper.New()

	v.Set("name", "gopher")

	fmt.Println("name:", v.Get("name"))
}

func Write() {
	v := viper.New()

	v.SetConfigName("config")
	v.SetConfigType("json")
	v.AddConfigPath(".")
	v.AddConfigPath("./06-libraries/02-config/01-viper/config")

	if err := v.ReadInConfig(); err != nil {
		fmt.Println("error:", err)
		return
	}

	v.Set("appName", "Here We Go!")

	if err := v.WriteConfig(); err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("appName:", v.Get("appName"))
}

func Keys() {
	v := viper.New()

	v.Set("name", "gopher")

	keys := v.AllKeys()

	for _, key := range keys {
		fmt.Println(key, ">", v.Get(key))
	}
}

func Settings() {
	v := viper.New()

	v.Set("name", "gopher")

	list := v.AllSettings()

	for key, value := range list {
		fmt.Println(key, ">", value)
	}
}

func Watch() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./06-libraries/02-config/01-viper/config")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("error:", err)
		return
	}

	viper.WatchConfig()

	fmt.Println("appName:", viper.Get("appName"))

	viper.Reset()
}

func Sample(name string, fn func())  {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}