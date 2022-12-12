package main

import (
	"testing"

	"github.com/kataras/iris/v12/httptest"
)

func TestMain(t *testing.T) {

	app := newApp()
	e := httptest.New(t, app)

	e.GET("/demo/test").WithBasicAuth("test", "test").Expect().Status(httptest.StatusOK)

}
