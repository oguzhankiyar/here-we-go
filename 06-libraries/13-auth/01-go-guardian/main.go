package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/shaj13/go-guardian/v2/auth"
	"github.com/shaj13/go-guardian/v2/auth/strategies/basic"
	"github.com/shaj13/go-guardian/v2/auth/strategies/jwt"
	"github.com/shaj13/go-guardian/v2/auth/strategies/union"
	"github.com/shaj13/libcache"
	_ "github.com/shaj13/libcache/fifo"
)

var strategy union.Union
var keeper jwt.SecretsKeeper

func main() {
	keeper = jwt.StaticSecret{
		ID:        "secret-id",
		Secret:    []byte("secret"),
		Algorithm: jwt.HS256,
	}

	cache := libcache.FIFO.New(0)
	cache.SetTTL(time.Minute * 5)
	cache.RegisterOnExpired(func(key, _ interface{}) {
		cache.Peek(key)
	})

	basicStrategy := basic.NewCached(ValidateUser, cache)
	jwtStrategy := jwt.New(cache, keeper)
	strategy = union.New(jwtStrategy, basicStrategy)

	router := mux.NewRouter()
	router.HandleFunc("/v1/auth/token", Middleware(http.HandlerFunc(CreateToken))).Methods("GET")
	router.HandleFunc("/v1/books", Middleware(http.HandlerFunc(GetBookAuthor))).Methods("GET")

	fmt.Println("listening on http://127.0.0.1:2805")
	fmt.Println("run 'curl -k http://127.0.0.1:2805/v1/auth/token -u admin:123456'")
	fmt.Println("run 'curl -k http://127.0.0.1:2805/v1/books -H \"Authorization: Bearer <TOKEN>\"'")

	http.ListenAndServe("127.0.0.1:2805", router)
}

func CreateToken(w http.ResponseWriter, r *http.Request) {
	user := auth.User(r)
	token, _ := jwt.IssueAccessToken(user, keeper)
	body, err := json.Marshal(map[string]interface{}{
		"token": token,
	})
	if err != nil {
		panic(err)
	}
	w.Write(body)
}

func GetBookAuthor(w http.ResponseWriter, r *http.Request) {
	books := []map[string]interface{}{
		{
			"id": 100,
			"name": "The Golang",
		},
		{
			"id": 200,
			"name": "The Gopher",
		},
		{
			"id": 300,
			"name": "Computer Science",
		},
	}

	body, err := json.Marshal(books)
	if err != nil {
		panic(err)
	}
	w.Write(body)
}

func ValidateUser(ctx context.Context, r *http.Request, userName, password string) (auth.Info, error) {
	// Check from db
	if userName == "admin" && password == "123456" {
		return auth.NewDefaultUser("admin", "1", nil, nil), nil
	}

	return nil, fmt.Errorf("invalid user")
}

func Middleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, user, err := strategy.AuthenticateRequest(r)
		if err != nil {
			fmt.Println(err)
			code := http.StatusUnauthorized
			http.Error(w, http.StatusText(code), code)
			return
		}

		r = auth.RequestWithUser(user, r)
		next.ServeHTTP(w, r)
	}
}