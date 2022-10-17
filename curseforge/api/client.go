package api

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"
)

type CurseforgeClientHeader map[string][]string

func (header CurseforgeClientHeader) Add(additionalClientHeader *CurseforgeClientHeader) CurseforgeClientHeader {
	var newHeader CurseforgeClientHeader
	additionalClientHeaderEncoded, _ := json.Marshal(additionalClientHeader)
	headerMarshaled, _ := json.Marshal(header)

	_ = json.Unmarshal(headerMarshaled, &newHeader)
	_ = json.Unmarshal(headerMarshaled, &additionalClientHeaderEncoded)

	return newHeader
}

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

// newCurseforgeClientForRoute returns a new client that calls uri.
func (api CurseforgeApi) newCurseforgeClientForRoute(uri string) *CurseforgeClient {
	return &CurseforgeClient{uri: uri, method: http.MethodGet, headers: api.clientHeader(), baseUrl: api.baseURL()}
}

func (client *CurseforgeClient) BaseUrl(baseUrl url.URL) {
	client.baseUrl = baseUrl
}

func (client *CurseforgeClient) Method(method string) {
	client.method = method
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
