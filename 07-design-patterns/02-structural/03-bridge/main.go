package main

import "fmt"

type Client interface {
	Request(string) string
}

type AuthClient struct {
	key string
}

func (c AuthClient) Request(url string) string {
	return fmt.Sprintf("response for %s with %s", url, c.key)
}

type NoAuthClient struct {

}

func (c NoAuthClient) Request(url string) string {
	return fmt.Sprintf("response for %s", url)
}

type Requester struct {
	client Client
}

func (r Requester) Request(url string) string {
	return r.client.Request(url)
}

func main() {
	authClient := AuthClient{"secure_key"}
	noAuthClient := NoAuthClient{}

	authRequester := Requester{authClient}
	noAuthRequester := Requester{noAuthClient}

	fmt.Printf("authRequest -> %s\n", authRequester.Request("localhost"))
	fmt.Printf("noAuthRequest -> %s\n", noAuthRequester.Request("localhost"))
}