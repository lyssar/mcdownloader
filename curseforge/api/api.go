package api

import (
	"net/url"
)

type CurseforgeApi struct {
	Config CurseforgeApiConfig
}

type CurseforgeApiConfig struct {
	BaseUrlProtocol  string
	BaseUrl          string
	ApiKey           string
	AdditionalHeader *CurseforgeClientHeader
}

func (api CurseforgeApi) clientHeader() CurseforgeClientHeader {
	defaults := CurseforgeClientHeader{
		"Content-Type": []string{"application/json"},
		"Accept":       []string{"application/json"},
		"x-api-key":    []string{api.Config.ApiKey},
	}
	if api.Config.AdditionalHeader != nil {
		return defaults.Add(api.Config.AdditionalHeader)
	}
	return defaults
}

func (api CurseforgeApi) baseURL() url.URL {
	return url.URL{
		Scheme: api.Config.BaseUrlProtocol,
		Host:   api.Config.BaseUrl,
		Path:   "",
	}
}
