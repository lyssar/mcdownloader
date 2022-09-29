package api

import (
	"encoding/json"
	"github.com/lyssar/msdcli/utils"
	"io"
	"net/http"
	"net/url"
	"time"
)

type CurseforgeClientHeader map[string][]string

// CurseforgeClient Curseforge api client
type CurseforgeClient struct {
	headers     CurseforgeClientHeader
	uri         string
	method      string
	baseUrl     url.URL
	queryValues *url.Values
	bodyData    *io.Reader
	response    *http.Response
	requestErr  error
}

var defaultBaseURL = url.URL{
	Scheme: utils.GetConfig().CurseForge.BaseUrlProtocol,
	Host:   utils.GetConfig().CurseForge.BaseUrl,
	Path:   "",
}

var defaultHeader = CurseforgeClientHeader{
	"Accept":    []string{"application/json"},
	"x-api-key": []string{utils.GetConfig().CurseForge.ApiKey},
}

func createHttpRequest(method string, url *url.URL, queryValues *url.Values, data *io.Reader, headers CurseforgeClientHeader) (req *http.Request, err error) {
	if data == nil {
		data = new(io.Reader)
	}
	req, err = http.NewRequest(method, url.String(), *data)
	req.Header = http.Header(headers)
	if queryValues != nil {
		req.URL.RawQuery = queryValues.Encode()
	}

	return
}

// NewCurseforgeClientForRoute returns a new client that calls uri.
func NewCurseforgeClientForRoute(uri string) *CurseforgeClient {
	return &CurseforgeClient{uri: uri, method: http.MethodGet, headers: defaultHeader, baseUrl: defaultBaseURL}
}

func (client *CurseforgeClient) BaseUrl(baseUrl url.URL) {
	client.baseUrl = baseUrl
}

func (client *CurseforgeClient) Method(method string) {
	client.method = method
}

func (client *CurseforgeClient) Query(queryValues *url.Values) {
	client.queryValues = queryValues
}

func (client *CurseforgeClient) Data(bodyData io.Reader) {
	client.bodyData = &bodyData
}

func (client *CurseforgeClient) Request() {
	endpoint := client.baseUrl.ResolveReference(&url.URL{Path: client.uri})

	req, err := createHttpRequest(client.method, endpoint, client.queryValues, client.bodyData, client.headers)
	client.requestErr = err

	httpClient := &http.Client{Timeout: time.Minute}
	resp, err := httpClient.Do(req)
	client.requestErr = err
	client.response = resp
}

func (client *CurseforgeClient) GetContent(target interface{}) error {
	defer client.response.Body.Close()

	return json.NewDecoder(client.response.Body).Decode(target)
}
