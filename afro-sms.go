/**
 * Author: Amanuel Abay
 * File: afro-sms.go
 */

package afro

const (
	Version = "1.0.0"
)

// afroSmsOptions for CreateRequest
type afroSmsOptions struct {
	Key      string
	Endpoint string
	Host     string
	Subuser  string
}

type requestOptions struct {
	Auth     string
	Endpoint string
	Host     string
	Subuser  string
}

type Method string

// Supported HTTP verbs.
const (
	Get    Method = "GET"
	Post   Method = "POST"
	Put    Method = "PUT"
	Patch  Method = "PATCH"
	Delete Method = "DELETE"
)

type Request struct {
	Method      Method
	BaseURL     string // e.g. https://api.afromessage.com
	Headers     map[string]string
	QueryParams map[string]string
}

// add sender name to query param
func (r *Request) Sender(SenderName string) {
	r.QueryParams["sender"] = SenderName
}

// add receiver phone number to query param
func (r *Request) To(PhoneNumber string, Message string) {
	r.QueryParams["to"] = PhoneNumber
	r.QueryParams["message"] = Message
}

// add callback to query param
func (r *Request) CallBackUrl(CallBackUrl string) {
	r.QueryParams["callback"] = CallBackUrl
}

// return base url of afro sms
func (o *requestOptions) baseURL() string {
	return o.Host + o.Endpoint
}

// get request content
func GetRequest(key, endpoint, host string) Request {
	return createAfroSmsRequest(afroSmsOptions{key, "/api/send", host, ""})
}

// create afro sms request
func createAfroSmsRequest(afroOptions afroSmsOptions) Request {

	options := requestOptions{
		"Bearer " + afroOptions.Key,
		afroOptions.Endpoint,
		afroOptions.Host,
		afroOptions.Subuser,
	}

	if options.Host == "" {
		options.Host = "https://api.afromessage.com"
	}
	return requestNew(options)
}

// build request header and baseUrl
func requestNew(requestOptions requestOptions) Request {
	requestHeaders := map[string]string{
		"Authorization": requestOptions.Auth,
		"User-Agent":    "afromessage/" + Version + ";go",
		"Accept":        "application/json",
	}

	if len(requestOptions.Subuser) != 0 {
		requestHeaders["On-Behalf-Of"] = requestOptions.Subuser
	}

	return Request{
		BaseURL:     requestOptions.baseURL(),
		Headers:     requestHeaders,
		QueryParams: map[string]string{},
	}
}
