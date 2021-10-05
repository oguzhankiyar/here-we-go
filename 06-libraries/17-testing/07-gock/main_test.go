package main

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func Test(t *testing.T) {
	defer gock.Off()

	gock.New("http://foo.com").
		Get("/bar").
		Reply(200).
		JSON(map[string]string{"foo": "bar"})

	res, err := http.Get("http://foo.com/bar")

	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)

	body, _ := ioutil.ReadAll(res.Body)
	assert.Equal(t, "{\"foo\":\"bar\"}", string(body)[:13])

	assert.True(t, gock.IsDone())
}