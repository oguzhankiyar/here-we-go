package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

var waitGroup sync.WaitGroup
var server http.Server

func main() {
	Sample("ListenAndServe", ListenAndServe)
	Sample("Get", Get)
	Sample("Post", Post)
	Sample("Put", Put)
	Sample("Delete", Delete)
	Sample("Shutdown", Shutdown)
}

func ListenAndServe() {
	waitGroup = sync.WaitGroup{}
	waitGroup.Add(1)

	server = http.Server{Addr: ":8080"}

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Method: ")
		io.WriteString(w, req.Method)

		io.WriteString(w, " | ")
		io.WriteString(w, "Url: ")
		io.WriteString(w, req.URL.String())

		if req.Method != http.MethodPost && req.Method != http.MethodPut {
			return
		}

		io.WriteString(w, " | ")
		io.WriteString(w, "Data: ")

		body, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Println("error:", err)
			return
		}

		io.WriteString(w, string(body))
	}

	http.HandleFunc("/", helloHandler)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			fmt.Println("error:", err)
		}

		waitGroup.Done()
	}()
}

func Get() {
	res, err := http.Get("http://127.0.0.1:8080/list")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	data, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("%s\n", data)
}

func Post() {
	payload := make(map[string]string)
	payload["name"] = "Gopher"

	jsonBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	reader := bytes.NewReader(jsonBytes)

	res, err := http.Post("http://127.0.0.1:8080/create", "application/json", reader)
	if err != nil {
		fmt.Println("error:", err)
	}

	data, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("%s\n", data)
}

func Put() {
	payload := make(map[string]string)
	payload["name"] = "Go"

	jsonBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("error:", err)
	}

	reader := bytes.NewReader(jsonBytes)

	req, err := http.NewRequest(http.MethodPut, "http://127.0.0.1:8080/edit", reader)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := &http.Client{}

	res, err := client.Do(req)

	data, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("%s\n", data)
}

func Delete() {
	req, err := http.NewRequest(http.MethodDelete, "http://127.0.0.1:8080/delete", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s\n", data)
}

func Shutdown() {
	server.Shutdown(context.Background())
	waitGroup.Wait()
	fmt.Println("Shutdown completed")
}

func Sample(name string, fn func()) {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}