package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
)

func main() {
	Sample("Simple", Simple)
	Sample("Get", Get)
	Sample("Post", Post)
	Sample("Put", Put)
	Sample("Delete", Delete)
	Sample("Retry", Retry)
}

func Simple() {
	client := resty.New()

	resp, err := client.R().
		EnableTrace().
		Get("https://jsonplaceholder.typicode.com/todos/1")

	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()

	fmt.Println("Request Trace Info:")
	ti := resp.Request.TraceInfo()
	fmt.Println("  DNSLookup     :", ti.DNSLookup)
	fmt.Println("  ConnTime      :", ti.ConnTime)
	fmt.Println("  TCPConnTime   :", ti.TCPConnTime)
	fmt.Println("  TLSHandshake  :", ti.TLSHandshake)
	fmt.Println("  ServerTime    :", ti.ServerTime)
	fmt.Println("  ResponseTime  :", ti.ResponseTime)
	fmt.Println("  TotalTime     :", ti.TotalTime)
	fmt.Println("  IsConnReused  :", ti.IsConnReused)
	fmt.Println("  IsConnWasIdle :", ti.IsConnWasIdle)
	fmt.Println("  ConnIdleTime  :", ti.ConnIdleTime)
	fmt.Println("  RequestAttempt:", ti.RequestAttempt)
	fmt.Println("  RemoteAddr    :", ti.RemoteAddr.String())
}

func Get() {
	client := resty.New()

	resp, err := client.R().
		SetQueryParams(map[string]string{
			"page_no": "1",
			"limit": "20",
			"sort":"name",
			"order": "asc",
			"random":strconv.FormatInt(time.Now().Unix(), 10),
		}).
		SetHeader("Accept", "application/json").
		SetAuthToken("BC594900518B4F7EAC75BD37F019E08FBC594900518B4F7EAC75BD37F019E08F").
		Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println(resp)
}

func Post() {
	client := resty.New()

	type User struct {
		Username 	string		`json:"username"`
		Password 	string		`json:"password"`
	}

	type AuthSuccess struct {
		Token 		string		`json:"token"`
	}

	type AuthError struct {
		Message		string		`json:"token"`
	}

	resp, err := client.R().
		SetBody(User{Username: "admin", Password: "123456"}).
		SetResult(&AuthSuccess{}).
		SetError(&AuthError{}).
		Post("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println(resp)
}

func Put() {
	client := resty.New()

	type Error struct {
		Message		string		`json:"message"`
	}

	resp, err := client.R().
		SetBody([]byte(`{"title":"The Golang", "author":"gopher"}`)).
		SetAuthToken("C6A79608-782F-4ED0-A11D-BD82FAD829CD").
		SetError(&Error{}).
		Put("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println(resp)
}

func Delete() {
	client := resty.New()

	type Error struct {
		Message		string		`json:"message"`
	}

	resp, err := client.R().
		SetAuthToken("C6A79608-782F-4ED0-A11D-BD82FAD829CD").
		SetError(&Error{}).
		Delete("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println(resp)
}

func Middleware() {
	client := resty.New()

	client.OnBeforeRequest(func(c *resty.Client, req *resty.Request) error {
		fmt.Println("request:", req)

		return nil
	})

	client.OnAfterResponse(func(c *resty.Client, resp *resty.Response) error {
		fmt.Println("response:", resp)

		return nil
	})

	client.OnError(func(req *resty.Request, err error) {
		if v, ok := err.(*resty.ResponseError); ok {
			fmt.Println("error:", v)
		}
	})

	resp, err := client.R().
		Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println(resp)
}

func Retry() {
	client := resty.New()

	client.
		SetRetryCount(2).
		SetRetryWaitTime(500 * time.Millisecond).
		SetRetryMaxWaitTime(20 * time.Second).
		AddRetryHook(func(response *resty.Response, err error) {
			fmt.Println(time.Now(), "retrying")
		}).
		AddRetryCondition(
			func(r *resty.Response, err error) bool {
				return !r.IsSuccess()
			},
		)

	resp, err := client.R().
		Get("https://run.mocky.io/v3/99c7664d-5ed1-4ff2-b453-fbb0b4b54c4c")

	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println(resp)
}

func Sample(name string, fn func())  {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}