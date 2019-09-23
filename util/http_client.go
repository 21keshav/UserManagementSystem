package util

import (
	"io"
	"net/http"
)

const GET = "GET"
const POST = "POST"
const PATCH = "PATCH"
const DELETE = "GET"
const PUT = "PUT"

//go:generate counterfeiter -o fakes/fake_http_client.go --fake-name FakeHttpClient . HttpClientUtil
type HttpClientUtil interface {
	Get(url string) (*http.Response, error)
	Post(url, contentType string, body io.Reader) (*http.Response, error)
	Patch(url string, body io.Reader) (*http.Response, error)
	PUT(url string, body io.Reader) (*http.Response, error)
}

type HttpClientUtilImpl struct {
	httpClient http.Client
}

func NewHttpClientUtil(options HttpOptions) HttpClientUtil {

	client := http.Client{
		Timeout: options.timeout,
	}
	return &HttpClientUtilImpl{
		httpClient: client,
	}

}

func (h *HttpClientUtilImpl) Get(url string) (*http.Response, error) {
	resp, err := h.doHttpRequest(GET, url, nil)
	return resp, err
}

func (h *HttpClientUtilImpl) Post(url, contentType string, body io.Reader) (*http.Response, error) {
	resp, err := h.doHttpRequest(POST, url, body)
	return resp, err
}

func (h *HttpClientUtilImpl) Delete(url string, body io.Reader) (*http.Response, error) {
	resp, err := h.doHttpRequest(DELETE, url, body)
	return resp, err
}
func (h *HttpClientUtilImpl) Patch(url string, body io.Reader) (*http.Response, error) {
	resp, err := h.doHttpRequest(DELETE, url, body)
	return resp, err
}
func (h *HttpClientUtilImpl) PUT(url string, body io.Reader) (*http.Response, error) {
	resp, err := h.doHttpRequest(PUT, url, body)
	return resp, err
}

func (h *HttpClientUtilImpl) doHttpRequest(HttpRequestType, url string, body io.Reader) (*http.Response, error) {
	var (
		err  error
		req  *http.Request
		resp *http.Response
	)

	if req, err = http.NewRequest(HttpRequestType, url, body); err != nil {
		return nil, err
	}

	if resp, err = h.httpClient.Do(req); err != nil {
		return nil, err
	}
	return resp, nil
}
