/**
 * Author: Amanuel Abay
 * File: api.go
 */

package afro

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
)

// network response type
type Response map[string]interface{}

// make request
func (c *Client) MakeRequest(req *http.Request) (*http.Response, error) {
	return c.HTTPClient.Do(req)
}

// http client type
type Client struct {
	HTTPClient *http.Client
}

// declare http client
var DefaultClient = &Client{HTTPClient: &http.Client{}}

// build request to sent request
func Api(request Request) (Response, error) {
	return MakeRequest(request)
}

// MakeRequest make the request with context
func MakeRequest(request Request) (Response, error) {
	return MakeRequestWithContext(context.Background(), request)
}

func MakeRequestWithContext(ctx context.Context, request Request) (Response, error) {
	return DefaultClient.SendWithContext(ctx, request)
}

// AddQueryParameters adds query parameters to the URL.
func AddQueryParameters(baseURL string, queryParams map[string]string) string {
	baseURL += "?"
	params := url.Values{}
	for key, value := range queryParams {
		params.Add(key, value)
	}
	return baseURL + params.Encode()
}

// BuildRequestObject creates the HTTP request object.
func BuildRequestObject(request Request) (*http.Request, error) {

	// Add any query parameters to the URL.
	if len(request.QueryParams) != 0 {
		request.BaseURL = AddQueryParameters(request.BaseURL, request.QueryParams)
	}
	req, err := http.NewRequest(string(request.Method), request.BaseURL, nil)
	if err != nil {
		return req, err
	}

	// get key and value for request header,
	for key, value := range request.Headers {
		req.Header.Set(key, value)
	}
	return req, err
}

// BuildResponse builds the response struct.
func BuildResponse(res *http.Response) (response Response, err error) {

	// get and decode request response
	err = json.NewDecoder(res.Body).Decode(&response)

	if err != nil {
		return nil, err
	}

	return response, err
}

// SendWithContext will build your request passing in the provided context, make the request, and build your response.
func (c *Client) SendWithContext(ctx context.Context, request Request) (Response, error) {

	// Build the HTTP request object.
	req, err := BuildRequestObject(request)
	if err != nil {
		return nil, err
	}
	// Pass in the user provided context
	req = req.WithContext(ctx)

	// Build the HTTP client and make the request.
	res, err := c.MakeRequest(req)
	if err != nil {
		return nil, err
	}

	// close request body after the request has been done successfully
	defer res.Body.Close()

	// Build Response object.
	return BuildResponse(res)
}
