package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gojek/heimdall/v7/httpclient"
	"github.com/gojek/heimdall/v7/plugins"
)

func main() {
	Sample("Do", Do)
	Sample("Get", Get)
	Sample("Post", Post)
	Sample("Put", Put)
	Sample("Delete", Delete)
	Sample("Logger", Logger)
}

func Do() {
	timeout := 1000 * time.Millisecond
	client := httpclient.NewClient(httpclient.WithHTTPTimeout(timeout))

	req, _ := http.NewRequest(http.MethodGet, "https://run.mocky.io/v3/41ac0f88-ff00-4fb5-a7b7-a749781c3260", nil)
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}

func Get() {
	timeout := 1000 * time.Millisecond
	client := httpclient.NewClient(httpclient.WithHTTPTimeout(timeout))

	res, err := client.Get("https://run.mocky.io/v3/41ac0f88-ff00-4fb5-a7b7-a749781c3260", nil)
	if err != nil{
		panic(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}

func Post() {
	timeout := 1000 * time.Millisecond
	client := httpclient.NewClient(httpclient.WithHTTPTimeout(timeout))

	payload := bytes.NewReader([]byte("{ \"name\": \"gopher\" }"))

	res, err := client.Post("https://run.mocky.io/v3/41ac0f88-ff00-4fb5-a7b7-a749781c3260", payload, nil)
	if err != nil{
		panic(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}

func Put() {
	timeout := 1000 * time.Millisecond
	client := httpclient.NewClient(httpclient.WithHTTPTimeout(timeout))

	payload := bytes.NewReader([]byte("{ \"name\": \"gopher\" }"))

	res, err := client.Put("https://run.mocky.io/v3/41ac0f88-ff00-4fb5-a7b7-a749781c3260", payload, nil)
	if err != nil{
		panic(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}

func Delete() {
	timeout := 1000 * time.Millisecond
	client := httpclient.NewClient(httpclient.WithHTTPTimeout(timeout))

	res, err := client.Delete("https://run.mocky.io/v3/41ac0f88-ff00-4fb5-a7b7-a749781c3260", nil)
	if err != nil{
		panic(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}

func Logger() {
	timeout := 1000 * time.Millisecond
	client := httpclient.NewClient(httpclient.WithHTTPTimeout(timeout))
	requestLogger := plugins.NewRequestLogger(nil, nil)
	client.AddPlugin(requestLogger)

	req, _ := http.NewRequest(http.MethodGet, "https://run.mocky.io/v3/41ac0f88-ff00-4fb5-a7b7-a749781c3260", nil)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("error:", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}

func Sample(name string, fn func())  {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}