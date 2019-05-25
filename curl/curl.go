package curl

import (
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

type Response struct {
	Raw     *http.Response
	Headers map[string]string
	Body    string
}

func NewResponse() *Response {
	return &Response{}
}

func (this *Response) IsOk() bool {
	return this.Raw.StatusCode == 200
}

func (this *Response) parseHeaders() error {
	headers := map[string]string{}
	for k, v := range this.Raw.Header {
		headers[k] = v[0]
	}
	this.Headers = headers
	return nil
}

func (this *Response) parseBody() error {
	if body, err := ioutil.ReadAll(this.Raw.Body); err != nil {
		panic(err)
	} else {
		this.Body = string(body)
	}
	return nil
}

var (
	GET_METHOD    = "GET"
	POST_METHOD   = "POST"
	SENDTYPE_FORM = "form"
	SENDTYPE_JSON = "json"
)

type Request struct {
	client          *http.Client
	req             *http.Request
	Method          string
	Url             string
	dialTimeout     time.Duration
	responseTimeOut time.Duration
	Headers         map[string]string
	Cookies         map[string]string
	Queries         map[string]string
	Body            map[string]interface{}
	SendType        string
}

func NewRequest() *Request {
	r := &Request{}
	r.dialTimeout = 15
	r.responseTimeOut = 15
	r.Headers = map[string]string{}
	r.Cookies = map[string]string{}
	r.Queries = map[string]string{}
	return r
}

func (this *Request) SetMethod(method string) *Request {
	this.Method = method
	return this
}

func (this *Request) SetUrl(url string) *Request {
	this.Url = url
	return this
}

func (this *Request) AddHeaders(headers map[string]string) *Request {
	for k, v := range headers {
		this.Headers[k] = v
	}
	return this
}

func (this *Request) AddHeader(k string, v string) *Request {
	this.Headers[k] = v
	return this
}

func (this *Request) setHeaders() {
	for k, v := range this.Headers {
		this.req.Header.Set(k, v)
	}
}

func (this *Request) AddFormHeader() *Request {
	this.AddHeader("Content-Type", "application/x-www-form-urlencoded; charset=utf-8;")
	return this
}

func (this *Request) AddFormDataHeader() *Request {
	this.AddHeader("Content-Type", "multipart/form-data; charset=utf-8;")
	return this
}

func (this *Request) AddJsonHeader() *Request {
	this.AddHeader("Content-Type", "application/json")
	this.SendType = SENDTYPE_JSON
	return this
}

func (this *Request) AddPlainHeader() *Request {
	this.AddHeader("Content-Type", "application/plain")
	return this
}

func (this *Request) AddCookies(cookies map[string]string) *Request {
	for k, v := range cookies {
		this.Cookies[k] = v
	}
	return this
}

func (this *Request) setCookies() {
	for k, v := range this.Cookies {
		this.req.AddCookie(&http.Cookie{
			Name:  k,
			Value: v,
		})
	}
}

func (this *Request) AddQueries(queries map[string]string) *Request {
	for k, v := range queries {
		this.Queries[k] = v
	}
	return this
}

func (this *Request) setQueries() {
	q := this.req.URL.Query()
	for k, v := range this.Queries {
		q.Add(k, v)
	}
	this.req.URL.RawQuery = q.Encode()
}

func (this *Request) SetDialTimeOut(TimeOutSecond int) {
	this.dialTimeout = time.Duration(TimeOutSecond)
}

func (this *Request) SetResponseTimeOut(TimeOutSecond int) {
	this.responseTimeOut = time.Duration(TimeOutSecond)
}

func (this *Request) SetBody(Body map[string]interface{}) *Request {
	this.Body = Body
	return this
}

func (this *Request) Form() (*Response, error) {
	return this.AddFormHeader().Send(this.Url, http.MethodPost)
}

func (this *Request) FormData() (*Response, error) {
	return this.AddFormDataHeader().Send(this.Url, http.MethodPost)
}

func (this *Request) Json() (*Response, error) {
	return this.AddJsonHeader().Send(this.Url, http.MethodPost)
}

func (this *Request) Get() (*Response, error) {
	return this.Send(this.Url, http.MethodGet)
}

func (this *Request) Post() (*Response, error) {
	return this.Send(this.Url, http.MethodPost)
}

func (this *Request) Put() (*Response, error) {
	return this.AddFormHeader().Send(this.Url, http.MethodPut)
}

func (this *Request) Patch() (*Response, error) {
	return this.AddFormHeader().Send(this.Url, http.MethodPatch)
}

func (this *Request) Delete() (*Response, error) {
	return this.AddFormHeader().Send(this.Url, http.MethodDelete)
}

func (this *Request) NewClient() {
	this.client = &http.Client{
		Transport: &http.Transport{
			Dial: func(network, addr string) (net.Conn, error) {
				conn, err := net.DialTimeout(network, addr, time.Second*this.dialTimeout)
				if err != nil {
					return nil, err
				}
				conn.SetDeadline(time.Now().Add(time.Second * this.dialTimeout))
				return conn, nil
			},
			ResponseHeaderTimeout: time.Second * this.responseTimeOut,
		},
	}
}

func (this *Request) Send(url string, method string) (*Response, error) {
	this.NewClient()

	var body string
	if this.Body != nil {
		if this.SendType == SENDTYPE_JSON {
			jsonByte, err := json.Marshal(this.Body)
			if err != nil {
				return nil, err
			}
			body = string(jsonByte)
		} else {
			send_body := http.Request{}
			send_body.ParseForm()
			for k, v := range this.Body {
				send_body.Form.Add(k, v.(string))
			}
			body = send_body.Form.Encode()
		}
	}

	if req, err := http.NewRequest(method, url, strings.NewReader(body)); err != nil {
		return nil, err
	} else {
		this.req = req
	}
	this.setHeaders()
	this.setCookies()
	this.setQueries()

	response := NewResponse()
	if resp, err := this.client.Do(this.req); err != nil {
		return nil, err
	} else {
		response.Raw = resp
	}
	defer response.Raw.Body.Close()

	response.parseHeaders()
	response.parseBody()

	return response, nil
}
