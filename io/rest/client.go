package rest

import (
	"fmt"

	"github.com/cyrildever/go-utls/io/rest/header/content_type"
	"github.com/valyala/fasthttp"
)

//--- TYPES

// Client ...
type Client struct {
	// URL is the full URL (including protocol and eventual query string et al.)
	URL string

	// Headers is an optional map of name -> value tuples of HTTP headers
	Headers map[string]string

	// Context gives the name of the calling context, eg. "rest"
	Context string
}

//--- METHODS

// Get ...
func (c *Client) Get() (statusCode int, body []byte, err error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.SetRequestURI(c.URL)
	req.Header.SetMethod(fasthttp.MethodGet)

	if len(c.Headers) > 0 {
		for k, v := range c.Headers {
			req.Header.Add(k, v)
		}
	}

	e := fasthttp.Do(req, resp)
	if e != nil {
		err = fmt.Errorf("%s responded with status %d and error message: %s [%s] - %s", c.Context, resp.StatusCode(), string(resp.Body()), e.Error(), c.URL)
		return
	}

	statusCode = resp.StatusCode()
	if len(resp.Body()) > 0 {
		var copyOfBody = make([]byte, len(resp.Body()))
		copy(copyOfBody, resp.Body())
		body = copyOfBody
	}
	return
}

// Post ...
func (c *Client) Post(data []byte, contentType string) (statusCode int, body []byte, err error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.SetRequestURI(c.URL)
	req.SetBody(data)
	req.Header.SetMethod(fasthttp.MethodPost)
	if contentType != "" {
		if !content_type.IsAuthorized(contentType) {
			err = fmt.Errorf("unauthorized content type: %s", contentType)
			return
		}
		req.Header.SetContentType(contentType)
	}
	if len(c.Headers) > 0 {
		for k, v := range c.Headers {
			req.Header.Add(k, v)
		}
	}

	e := fasthttp.Do(req, resp)
	if e != nil {
		err = fmt.Errorf("%s responded with status %d and error message: %s [%s] - %s", c.Context, resp.StatusCode(), string(resp.Body()), e.Error(), c.URL)
		return
	}

	if len(resp.Body()) > 0 {
		var copyOfBody = make([]byte, len(resp.Body()))
		copy(copyOfBody, resp.Body())
		body = copyOfBody
	}
	statusCode = resp.StatusCode()
	return
}
