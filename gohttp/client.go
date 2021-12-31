package gohttp

import (
	"net/http"
)

type httpClient struct {
	Headers http.Header
}

func New() HttpClient {
	client := &httpClient{} // & means get the memory address of the variable, httpClient{} instatiates the struct
	return client           // return a pointer to the instantiated struct, otherwise we just get a copy
}

// article about the interface type here: https://www.alexedwards.net/blog/interfaces-explained
// using an interface here means that we can easily create mocks of HttpClient in unit tests - all they have to do is implement the methods on the interface
type HttpClient interface { // this is publicly available from the module because of the capitals
	SetHeaders(headers http.Header)
	Get(url string, headers http.Header) (*http.Response, error)
	Post(url string, headers http.Header, body interface{}) (*http.Response, error)
	Put(url string, headers http.Header, body interface{}) (*http.Response, error)
	Patch(url string, headers http.Header, body interface{}) (*http.Response, error)
	Delete(url string, headers http.Header) (*http.Response, error)
}

func (c *httpClient) SetHeaders(headers http.Header) {
	c.Headers = headers
}

// these methods get assigned to the httpClient struct type automatically.
// because of these methods, the struct implements the HttpClient type automatically (don't need to say 'implements')
func (c *httpClient) Get(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodGet, url, headers, nil)
}

// the interface{} type here is used to denote any type (I think) - e.g. a struct type in this case
func (c *httpClient) Post(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPost, url, headers, body)
}

func (c *httpClient) Put(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPut, url, headers, body)
}

func (c *httpClient) Patch(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPatch, url, headers, body)
}

func (c *httpClient) Delete(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodDelete, url, headers, nil)
}
