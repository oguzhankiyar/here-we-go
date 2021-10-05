package main

import (
	"fmt"
	"log"
)

type Middleware interface {
	Pipe(interface{}) (interface{}, error)
}

type RequestLogger struct {
	Next Middleware
}

func (l RequestLogger) Pipe(data interface{}) (interface{}, error) {
	request := data.(string)
	log.Println("request:", request)

	if l.Next != nil {
		var err error
		data, err = l.Next.Pipe(data)
		if err != nil {
			log.Panicln(err)
		}
	}

	return data, nil
}

type ResponseLogger struct {
	Next Middleware
}

func (l ResponseLogger) Pipe(data interface{}) (interface{}, error) {
	request := data.(string)
	log.Println("response:", request)

	if l.Next != nil {
		var err error
		data, err = l.Next.Pipe(data)
		if err != nil {
			log.Panicln(err)
		}
	}

	return data, nil
}

type ExceptionHandler struct {
	Next Middleware
}

func (h ExceptionHandler) Pipe(data interface{}) (interface{}, error) {

	if h.Next != nil {
		var err error
		data, err = h.Next.Pipe(data)
		if err != nil {
			log.Panicln(err)
		}
	}

	return data, nil
}

type RequestProcessor struct {
	Next Middleware
}

func (p RequestProcessor) Pipe(data interface{}) (interface{}, error) {
	data = `{ "status": true }`

	if p.Next != nil {
		var err error
		data, err = p.Next.Pipe(data)
		if err != nil {
			log.Panicln(err)
		}
	}

	return data, nil
}

func main() {
	responseLogger := ResponseLogger{nil}
	requestProcessor := RequestProcessor{responseLogger}
	exceptionHandler := ExceptionHandler{requestProcessor}
	requestLogger := RequestLogger{exceptionHandler}

	_, err := requestLogger.Pipe(`{ "id": 1 }`)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
}