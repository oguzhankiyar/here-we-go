package main

import (
	"fmt"
	"net/http"

	"github.com/openshift/osin"
)

func main() {
	config := osin.NewServerConfig()
	config.AllowedAuthorizeTypes = osin.AllowedAuthorizeType{osin.CODE, osin.TOKEN}
	config.AllowedAccessTypes = osin.AllowedAccessType{
		osin.AUTHORIZATION_CODE,
		osin.REFRESH_TOKEN,
		osin.PASSWORD,
		osin.CLIENT_CREDENTIALS,
		osin.ASSERTION,
	}
	config.AllowGetAccessRequest = true
	config.AllowClientSecretInParams = true
	server := osin.NewServer(config, NewTestStorage())

	app := App{server}

	// Application home endpoint
	http.HandleFunc("/", app.Home)

	// Authorization code endpoint
	http.HandleFunc("/authorize", app.Authorize)

	// Access token endpoint
	http.HandleFunc("/token", app.Token)

	// Information endpoint
	http.HandleFunc("/info", app.Info)

	// Application destination - CODE
	http.HandleFunc("/appauth/code", app.AuthCode)

	// Application destination - TOKEN
	http.HandleFunc("/appauth/token", app.AuthToken)

	// Application destination - PASSWORD
	http.HandleFunc("/appauth/password", app.AuthPassword)

	// Application destination - CLIENT_CREDENTIALS
	http.HandleFunc("/appauth/client_credentials", app.AuthCredentials)

	// Application destination - ASSERTION
	http.HandleFunc("/appauth/assertion", app.AuthAssertion)

	// Application destination - REFRESH
	http.HandleFunc("/appauth/refresh", app.AuthRefresh)

	// Application destination - INFO
	http.HandleFunc("/appauth/info", app.AuthInfo)

	fmt.Println("listening on http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}