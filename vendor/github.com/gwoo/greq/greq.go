// greq - simple http request library
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).

package greq

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Req object
type Req struct {
	host   string
	json   bool
	header http.Header
}

// Get a new req object.
func New(host string, json bool) *Req {
	return &Req{host, json, http.Header{}}
}

func (r *Req) Header(h http.Header) http.Header {
	if h != nil {
		r.header = h
	}
	return r.header
}

// Send head to the host.
func (r *Req) Head(path string) ([]byte, *http.Response, error) {
	return Do("HEAD", r.host+path, r.header, nil)
}

// Get the options from the host.
func (r *Req) Options(path string) ([]byte, *http.Response, error) {
	return Do("OPTIONS", r.host+path, r.header, nil)
}

// Get the path from the host.
func (r *Req) Get(path string) ([]byte, *http.Response, error) {
	return Do("GET", r.host+path, r.header, nil)
}

// Post data to the path on the host.
func (r *Req) Post(path string, data map[string]interface{}) ([]byte, *http.Response, error) {
	b, err := r.body(data)
	if err != nil {
		return nil, nil, err
	}
	return Do("POST", r.host+path, r.header, b)
}

// Put data to the path on the host.
func (r *Req) Put(path string, data map[string]interface{}) ([]byte, *http.Response, error) {
	b, err := r.body(data)
	if err != nil {
		return nil, nil, err
	}
	return Do("PUT", r.host+path, r.header, b)
}

// Send delete to the path on the host.
func (r *Req) Delete(path string) ([]byte, *http.Response, error) {
	return Do("DELETE", r.host+path, r.header, nil)
}

// Create an io.Reader for the body.
func (r *Req) body(data map[string]interface{}) (io.Reader, error) {
	if r.json == true {
		j, err := json.Marshal(data)
		if err == nil {
			return bytes.NewBuffer(j), nil
		}
		return nil, err
	}
	d := ToForm(data)
	if d != nil {
		return strings.NewReader(d.Encode()), nil
	}
	return nil, errors.New("Data could not be converted to values.")
}

// Generic request method
func Do(method string, url string, headers http.Header, body io.Reader) ([]byte, *http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, nil, err
	}
	req.Header = headers
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	b, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, nil, err
	}
	return b, res, nil
}

// Conver a map[string]interface to url.Values
func ToForm(data map[string]interface{}) url.Values {
	values := url.Values{}
	for k, v := range data {
		values.Set(k, fmt.Sprintf("%v", v))
	}
	return values
}
