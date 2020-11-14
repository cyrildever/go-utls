package rest

import (
	"fmt"

	"github.com/cyrildever/go-utls/io/rest/header/content_type"
	"github.com/valyala/fasthttp"
)

//--- TYPES

// APIClient defines the contract for a REST API client, the `expectedHeaders` argument being the names of the expected response headers returned by any call
type APIClient interface {
	Delete(expectedHeaders ...string) (statusCode int, body []byte, headers map[string]string, err error)
	Get(expectedHeaders ...string) (statusCode int, body []byte, headers map[string]string, err error)
	Patch(expectedHeaders ...string) (statusCode int, body []byte, headers map[string]string, err error)
	Post(data []byte, contentType string, expectedHeaders ...string) (statusCode int, body []byte, headers map[string]string, err error)
	Put(data []byte, contentType string, expectedHeaders ...string) (statusCode int, body []byte, headers map[string]string, err error)
}

// Client ...
type Client struct {
	// URL is the full URL (including protocol and eventual query string et al.)
	URL string

	// Headers is an optional map of name -> value tuples of HTTP request headers
	Headers map[string]string

	// Context gives the name of the calling context, eg. "rest"
	Context string
}

//--- METHODS

// Delete uses the URL and/or the query string to define which resource to delete, eg.
//	DELETE http://www.example.com/account/123
//	DELETE http://www.example.com/account?id=123
func (c *Client) Delete(expectedHeaders ...string) (statusCode int, body []byte, headers map[string]string, err error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.SetRequestURI(c.URL)
	req.Header.SetMethod(fasthttp.MethodDelete)
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

	if len(expectedHeaders) > 0 {
		headers = make(map[string]string, len(expectedHeaders))
		for _, h := range expectedHeaders {
			headers[h] = string(resp.Header.Peek(h))
		}
	}
	return
}

// Get uses the URL and/or the query string to define what resource to get, eg.
//	GET http://www.example.com/account/123
//	GET http://www.example.com/account?id=123
func (c *Client) Get(expectedHeaders ...string) (statusCode int, body []byte, headers map[string]string, err error) {
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
		err = fmt.Errorf("%s responded to GET with status %d and error message: %s [%s] - %s", c.Context, resp.StatusCode(), string(resp.Body()), e.Error(), c.URL)
		return
	}

	statusCode = resp.StatusCode()

	if len(resp.Body()) > 0 {
		var copyOfBody = make([]byte, len(resp.Body()))
		copy(copyOfBody, resp.Body())
		body = copyOfBody
	}

	if len(expectedHeaders) > 0 {
		headers = make(map[string]string, len(expectedHeaders))
		for _, h := range expectedHeaders {
			headers[h] = string(resp.Header.Peek(h))
		}
	}
	return
}

// Patch uses the URL and/or the query string to define which specific resource to update, eg.
//	PATCH http://www.example.com/account/123/name/Doe
//	PATCH http://www.example.com/account?id=123&name=Doe
func (c *Client) Patch(expectedHeaders ...string) (statusCode int, body []byte, headers map[string]string, err error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.SetRequestURI(c.URL)
	req.Header.SetMethod(fasthttp.MethodPatch)
	if len(c.Headers) > 0 {
		for k, v := range c.Headers {
			req.Header.Add(k, v)
		}
	}

	e := fasthttp.Do(req, resp)
	if e != nil {
		err = fmt.Errorf("%s responded to PATCH with status %d and error message: %s [%s] - %s", c.Context, resp.StatusCode(), string(resp.Body()), e.Error(), c.URL)
		return
	}

	if len(resp.Body()) > 0 {
		var copyOfBody = make([]byte, len(resp.Body()))
		copy(copyOfBody, resp.Body())
		body = copyOfBody
	}
	statusCode = resp.StatusCode()

	if len(expectedHeaders) > 0 {
		headers = make(map[string]string, len(expectedHeaders))
		for _, h := range expectedHeaders {
			headers[h] = string(resp.Header.Peek(h))
		}
	}
	return
}

// Post uses the URL and/or the query string to define what type of resource to add
//	POST http://www.example.com/account
func (c *Client) Post(data []byte, contentType string, expectedHeaders ...string) (statusCode int, body []byte, headers map[string]string, err error) {
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
		err = fmt.Errorf("%s responded to POST with status %d and error message: %s [%s] - %s", c.Context, resp.StatusCode(), string(resp.Body()), e.Error(), c.URL)
		return
	}

	if len(resp.Body()) > 0 {
		var copyOfBody = make([]byte, len(resp.Body()))
		copy(copyOfBody, resp.Body())
		body = copyOfBody
	}
	statusCode = resp.StatusCode()

	if len(expectedHeaders) > 0 {
		headers = make(map[string]string, len(expectedHeaders))
		for _, h := range expectedHeaders {
			headers[h] = string(resp.Header.Peek(h))
		}
	}
	return
}

// Put uses the URL and/or the query string to define which resource to update, eg.
//	PUT http://www.example.com/account/123
//	PUT http://www.example.com/account?id=123
func (c *Client) Put(data []byte, contentType string, expectedHeaders ...string) (statusCode int, body []byte, headers map[string]string, err error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.SetRequestURI(c.URL)
	req.SetBody(data)
	req.Header.SetMethod(fasthttp.MethodPut)
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
		err = fmt.Errorf("%s responded to PUT with status %d and error message: %s [%s] - %s", c.Context, resp.StatusCode(), string(resp.Body()), e.Error(), c.URL)
		return
	}

	if len(resp.Body()) > 0 {
		var copyOfBody = make([]byte, len(resp.Body()))
		copy(copyOfBody, resp.Body())
		body = copyOfBody
	}
	statusCode = resp.StatusCode()

	if len(expectedHeaders) > 0 {
		headers = make(map[string]string, len(expectedHeaders))
		for _, h := range expectedHeaders {
			headers[h] = string(resp.Header.Peek(h))
		}
	}
	return
}
