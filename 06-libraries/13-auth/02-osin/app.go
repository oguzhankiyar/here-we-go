package main

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/openshift/osin"
)

type App struct {
	Server *osin.Server
}

func (app App) Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<html><body>"))

	w.Write([]byte(fmt.Sprintf("<a href=\"/authorize?response_type=code&client_id=1234&state=xyz&scope=everything&redirect_uri=%s\">Code</a><br/>", url.QueryEscape("http://localhost:2805/appauth/code"))))
	w.Write([]byte(fmt.Sprintf("<a href=\"/authorize?response_type=token&client_id=1234&state=xyz&scope=everything&redirect_uri=%s\">Implicit</a><br/>", url.QueryEscape("http://localhost:2805/appauth/token"))))
	w.Write([]byte(fmt.Sprintf("<a href=\"/appauth/password\">Password</a><br/>")))
	w.Write([]byte(fmt.Sprintf("<a href=\"/appauth/client_credentials\">Client Credentials</a><br/>")))
	w.Write([]byte(fmt.Sprintf("<a href=\"/appauth/assertion\">Assertion</a><br/>")))

	w.Write([]byte("</body></html>"))
}

func (app App) Authorize(w http.ResponseWriter, r *http.Request) {
	resp := app.Server.NewResponse()
	defer resp.Close()

	if ar := app.Server.HandleAuthorizeRequest(resp, r); ar != nil {
		if !HandleLoginPage(ar, w, r) {
			return
		}
		ar.UserData = struct{ Login string }{Login: "test"}
		ar.Authorized = true
		app.Server.FinishAuthorizeRequest(resp, r, ar)
	}
	if resp.IsError && resp.InternalError != nil {
		fmt.Printf("ERROR: %s\n", resp.InternalError)
	}
	if !resp.IsError {
		resp.Output["custom_parameter"] = 187723
	}
	osin.OutputJSON(resp, w, r)
}

func (app App) Token(w http.ResponseWriter, r *http.Request) {
	resp := app.Server.NewResponse()
	defer resp.Close()

	if ar := app.Server.HandleAccessRequest(resp, r); ar != nil {
		switch ar.Type {
		case osin.AUTHORIZATION_CODE:
			ar.Authorized = true
		case osin.REFRESH_TOKEN:
			ar.Authorized = true
		case osin.PASSWORD:
			if ar.Username == "test" && ar.Password == "test" {
				ar.Authorized = true
			}
		case osin.CLIENT_CREDENTIALS:
			ar.Authorized = true
		case osin.ASSERTION:
			if ar.AssertionType == "urn:osin.example.complete" && ar.Assertion == "osin.data" {
				ar.Authorized = true
			}
		}
		app.Server.FinishAccessRequest(resp, r, ar)
	}
	if resp.IsError && resp.InternalError != nil {
		fmt.Printf("ERROR: %s\n", resp.InternalError)
	}
	if !resp.IsError {
		resp.Output["custom_parameter"] = 19923
	}
	osin.OutputJSON(resp, w, r)
}

func (app App) Info(w http.ResponseWriter, r *http.Request) {
	resp := app.Server.NewResponse()
	defer resp.Close()

	if ir := app.Server.HandleInfoRequest(resp, r); ir != nil {
		app.Server.FinishInfoRequest(resp, r, ir)
	}
	osin.OutputJSON(resp, w, r)
}

func (app App) AuthCode(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	code := r.FormValue("code")

	w.Write([]byte("<html><body>"))
	w.Write([]byte("APP AUTH - CODE<br/>"))
	defer w.Write([]byte("</body></html>"))

	if code == "" {
		w.Write([]byte("Nothing to do"))
		return
	}

	jr := make(map[string]interface{})

	// build access code url
	aurl := fmt.Sprintf("/token?grant_type=authorization_code&client_id=1234&client_secret=aabbccdd&state=xyz&redirect_uri=%s&code=%s",
		url.QueryEscape("http://localhost:2805/appauth/code"), url.QueryEscape(code))

	// if parse, download and parse json
	if r.FormValue("doparse") == "1" {
		err := DownloadAccessToken(fmt.Sprintf("http://localhost:2805%s", aurl),
			&osin.BasicAuth{"1234", "aabbccdd"}, jr)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.Write([]byte("<br/>"))
		}
	}

	// show json error
	if erd, ok := jr["error"]; ok {
		w.Write([]byte(fmt.Sprintf("ERROR: %s<br/>\n", erd)))
	}

	// show json access token
	if at, ok := jr["access_token"]; ok {
		w.Write([]byte(fmt.Sprintf("ACCESS TOKEN: %s<br/>\n", at)))
	}

	w.Write([]byte(fmt.Sprintf("FULL RESULT: %+v<br/>\n", jr)))

	// output links
	w.Write([]byte(fmt.Sprintf("<a href=\"%s\">Goto Token URL</a><br/>", aurl)))

	cururl := *r.URL
	curq := cururl.Query()
	curq.Add("doparse", "1")
	cururl.RawQuery = curq.Encode()
	w.Write([]byte(fmt.Sprintf("<a href=\"%s\">Download Token</a><br/>", cururl.String())))

	if rt, ok := jr["refresh_token"]; ok {
		rurl := fmt.Sprintf("/appauth/refresh?code=%s", rt)
		w.Write([]byte(fmt.Sprintf("<a href=\"%s\">Refresh Token</a><br/>", rurl)))
	}

	if at, ok := jr["access_token"]; ok {
		rurl := fmt.Sprintf("/appauth/info?code=%s", at)
		w.Write([]byte(fmt.Sprintf("<a href=\"%s\">Info</a><br/>", rurl)))
	}
}

func (app App) AuthToken(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	w.Write([]byte("<html><body>"))
	w.Write([]byte("APP AUTH - TOKEN<br/>"))

	w.Write([]byte("Response data in fragment - not acessible via server - Nothing to do"))

	w.Write([]byte("</body></html>"))
}

func (app App) AuthPassword(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	w.Write([]byte("<html><body>"))
	w.Write([]byte("APP AUTH - PASSWORD<br/>"))

	jr := make(map[string]interface{})

	// build access code url
	aurl := fmt.Sprintf("/token?grant_type=password&scope=everything&username=%s&password=%s",
		"test", "test")

	// download token
	err := DownloadAccessToken(fmt.Sprintf("http://localhost:2805%s", aurl),
		&osin.BasicAuth{Username: "1234", Password: "aabbccdd"}, jr)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.Write([]byte("<br/>"))
	}

	// show json error
	if erd, ok := jr["error"]; ok {
		w.Write([]byte(fmt.Sprintf("ERROR: %s<br/>\n", erd)))
	}

	// show json access token
	if at, ok := jr["access_token"]; ok {
		w.Write([]byte(fmt.Sprintf("ACCESS TOKEN: %s<br/>\n", at)))
	}

	w.Write([]byte(fmt.Sprintf("FULL RESULT: %+v<br/>\n", jr)))

	if rt, ok := jr["refresh_token"]; ok {
		rurl := fmt.Sprintf("/appauth/refresh?code=%s", rt)
		w.Write([]byte(fmt.Sprintf("<a href=\"%s\">Refresh Token</a><br/>", rurl)))
	}

	if at, ok := jr["access_token"]; ok {
		rurl := fmt.Sprintf("/appauth/info?code=%s", at)
		w.Write([]byte(fmt.Sprintf("<a href=\"%s\">Info</a><br/>", rurl)))
	}

	w.Write([]byte("</body></html>"))
}

func (app App) AuthCredentials(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	w.Write([]byte("<html><body>"))
	w.Write([]byte("APP AUTH - CLIENT CREDENTIALS<br/>"))

	jr := make(map[string]interface{})

	// build access code url
	aurl := fmt.Sprintf("/token?grant_type=client_credentials")

	// download token
	err := DownloadAccessToken(fmt.Sprintf("http://localhost:2805%s", aurl),
		&osin.BasicAuth{Username: "1234", Password: "aabbccdd"}, jr)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.Write([]byte("<br/>"))
	}

	// show json error
	if erd, ok := jr["error"]; ok {
		w.Write([]byte(fmt.Sprintf("ERROR: %s<br/>\n", erd)))
	}

	// show json access token
	if at, ok := jr["access_token"]; ok {
		w.Write([]byte(fmt.Sprintf("ACCESS TOKEN: %s<br/>\n", at)))
	}

	w.Write([]byte(fmt.Sprintf("FULL RESULT: %+v<br/>\n", jr)))

	if rt, ok := jr["refresh_token"]; ok {
		rurl := fmt.Sprintf("/appauth/refresh?code=%s", rt)
		w.Write([]byte(fmt.Sprintf("<a href=\"%s\">Refresh Token</a><br/>", rurl)))
	}

	if at, ok := jr["access_token"]; ok {
		rurl := fmt.Sprintf("/appauth/info?code=%s", at)
		w.Write([]byte(fmt.Sprintf("<a href=\"%s\">Info</a><br/>", rurl)))
	}

	w.Write([]byte("</body></html>"))
}

func (app App) AuthAssertion(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	w.Write([]byte("<html><body>"))
	w.Write([]byte("APP AUTH - ASSERTION<br/>"))

	jr := make(map[string]interface{})

	// build access code url
	aurl := fmt.Sprintf("/token?grant_type=assertion&assertion_type=urn:osin.example.complete&assertion=osin.data")

	// download token
	err := DownloadAccessToken(fmt.Sprintf("http://localhost:2805%s", aurl),
		&osin.BasicAuth{Username: "1234", Password: "aabbccdd"}, jr)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.Write([]byte("<br/>"))
	}

	// show json error
	if erd, ok := jr["error"]; ok {
		w.Write([]byte(fmt.Sprintf("ERROR: %s<br/>\n", erd)))
	}

	// show json access token
	if at, ok := jr["access_token"]; ok {
		w.Write([]byte(fmt.Sprintf("ACCESS TOKEN: %s<br/>\n", at)))
	}

	w.Write([]byte(fmt.Sprintf("FULL RESULT: %+v<br/>\n", jr)))

	if rt, ok := jr["refresh_token"]; ok {
		rurl := fmt.Sprintf("/appauth/refresh?code=%s", rt)
		w.Write([]byte(fmt.Sprintf("<a href=\"%s\">Refresh Token</a><br/>", rurl)))
	}

	if at, ok := jr["access_token"]; ok {
		rurl := fmt.Sprintf("/appauth/info?code=%s", at)
		w.Write([]byte(fmt.Sprintf("<a href=\"%s\">Info</a><br/>", rurl)))
	}

	w.Write([]byte("</body></html>"))
}

func (app App) AuthRefresh(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	w.Write([]byte("<html><body>"))
	w.Write([]byte("APP AUTH - REFRESH<br/>"))
	defer w.Write([]byte("</body></html>"))

	code := r.FormValue("code")

	if code == "" {
		w.Write([]byte("Nothing to do"))
		return
	}

	jr := make(map[string]interface{})

	// build access code url
	aurl := fmt.Sprintf("/token?grant_type=refresh_token&refresh_token=%s", url.QueryEscape(code))

	// download token
	err := DownloadAccessToken(fmt.Sprintf("http://localhost:2805%s", aurl),
		&osin.BasicAuth{Username: "1234", Password: "aabbccdd"}, jr)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.Write([]byte("<br/>"))
	}

	// show json error
	if erd, ok := jr["error"]; ok {
		w.Write([]byte(fmt.Sprintf("ERROR: %s<br/>\n", erd)))
	}

	// show json access token
	if at, ok := jr["access_token"]; ok {
		w.Write([]byte(fmt.Sprintf("ACCESS TOKEN: %s<br/>\n", at)))
	}

	w.Write([]byte(fmt.Sprintf("FULL RESULT: %+v<br/>\n", jr)))

	if rt, ok := jr["refresh_token"]; ok {
		rurl := fmt.Sprintf("/appauth/refresh?code=%s", rt)
		w.Write([]byte(fmt.Sprintf("<a href=\"%s\">Refresh Token</a><br/>", rurl)))
	}

	if at, ok := jr["access_token"]; ok {
		rurl := fmt.Sprintf("/appauth/info?code=%s", at)
		w.Write([]byte(fmt.Sprintf("<a href=\"%s\">Info</a><br/>", rurl)))
	}
}

func (app App) AuthInfo(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	w.Write([]byte("<html><body>"))
	w.Write([]byte("APP AUTH - INFO<br/>"))
	defer w.Write([]byte("</body></html>"))

	code := r.FormValue("code")

	if code == "" {
		w.Write([]byte("Nothing to do"))
		return
	}

	jr := make(map[string]interface{})

	// build access code url
	aurl := fmt.Sprintf("/info?code=%s", url.QueryEscape(code))

	// download token
	err := DownloadAccessToken(fmt.Sprintf("http://localhost:2805%s", aurl),
		&osin.BasicAuth{Username: "1234", Password: "aabbccdd"}, jr)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.Write([]byte("<br/>"))
	}

	// show json error
	if erd, ok := jr["error"]; ok {
		w.Write([]byte(fmt.Sprintf("ERROR: %s<br/>\n", erd)))
	}

	// show json access token
	if at, ok := jr["access_token"]; ok {
		w.Write([]byte(fmt.Sprintf("ACCESS TOKEN: %s<br/>\n", at)))
	}

	w.Write([]byte(fmt.Sprintf("FULL RESULT: %+v<br/>\n", jr)))

	if rt, ok := jr["refresh_token"]; ok {
		rurl := fmt.Sprintf("/appauth/refresh?code=%s", rt)
		w.Write([]byte(fmt.Sprintf("<a href=\"%s\">Refresh Token</a><br/>", rurl)))
	}
}