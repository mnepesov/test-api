package nasa

import (
	"encoding/json"
	"errors"
	"fmt"
	
	"github.com/valyala/fasthttp"
)

var (
	ErrSendRequest   = errors.New("can't send request")
	ErrEmptyBody     = errors.New("empty response")
	ErrParseResponse = errors.New("response parse error")
)

type INasa interface {
	GetAPOD() (APOD, error)
}

type Nasa struct {
	HttpClient *fasthttp.Client
	ApiKey     string
}

func NewNasaClient(httpClient *fasthttp.Client, apiKey string) INasa {
	return &Nasa{
		HttpClient: httpClient,
		ApiKey:     apiKey,
	}
}

func (n *Nasa) GetAPOD() (APOD, error) {
	req, resp := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(resp)
	}()
	
	url := fmt.Sprintf("https://api.nasa.gov/planetary/apod?api_key=%s", n.ApiKey)
	prepareRequest(req, fasthttp.MethodGet, url, nil)
	err := n.HttpClient.Do(req, resp)
	if err != nil {
		return APOD{}, ErrSendRequest
	}
	
	if len(resp.Body()) == 0 {
		return APOD{}, ErrEmptyBody
	}
	
	body := resp.Body()
	var apod APOD
	if err := json.Unmarshal(body, &apod); err != nil {
		return APOD{}, ErrParseResponse
	}
	
	return apod, nil
}

func prepareRequest(req *fasthttp.Request, method string, url string, body []byte) {
	req.Header.Set("Connection", "keep-alive")
	req.Header.SetMethod(method)
	
	req.Header.Set("Content-Type", "application/json")
	req.SetRequestURI(url)
	if method == fasthttp.MethodPost {
		req.SetBody(body)
		req.Header.SetContentLength(len(req.Body()))
	}
}
