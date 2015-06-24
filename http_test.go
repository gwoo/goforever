// goforever - processes management
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).

package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gwoo/greq"
)

func TestListHandler(t *testing.T) {
	daemon.children = children{
		"test": &Process{Name: "test"},
	}
	body, _ := newTestResponse("GET", "/", nil)
	ex := fmt.Sprintf("%s", string([]byte(`["test"]`)))
	r := fmt.Sprintf("%s", string(body))
	if ex != r {
		t.Errorf("\nExpected = %v\nResult = %v\n", ex, r)
	}
}

func TestShowHandler(t *testing.T) {
	daemon.children = children{
		"test": &Process{Name: "test"},
	}
	body, _ := newTestResponse("GET", "/test", nil)
	e := []byte(`{"Name":"test","Command":"","Args":null,"Pidfile":"","Logfile":"","Errfile":"","Path":"","Respawn":0,"Delay":"","Ping":"","Pid":0,"Status":""}`)
	ex := fmt.Sprintf("%s", e)
	r := fmt.Sprintf("%s", body)
	if ex != r {
		t.Errorf("\nExpected = %v\nResult = %v\n", ex, r)
	}
}

func TestPostHandler(t *testing.T) {
	daemon.children = children{
		"test": &Process{Name: "test", Command: "/bin/echo", Args: []string{"woohoo"}},
	}
	body, _ := newTestResponse("POST", "/test", nil)
	e := []byte(`{"Name":"test","Command":"/bin/echo","Args":["woohoo"],"Pidfile":"","Logfile":"","Errfile":"","Path":"","Respawn":0,"Delay":"","Ping":"","Pid":0,"Status":"stopped"}`)
	ex := fmt.Sprintf("%s", e)
	r := fmt.Sprintf("%s", body)
	if ex != r {
		t.Errorf("\nExpected = %v\nResult = %v\n", ex, r)
	}
}

func newTestResponse(method string, path string, body io.Reader) ([]byte, *http.Response) {
	ts := httptest.NewServer(http.HandlerFunc(Handler))
	defer ts.Close()
	url := ts.URL + path
	b, r, _ := greq.Do(method, url, nil, body)
	return b, r
}
