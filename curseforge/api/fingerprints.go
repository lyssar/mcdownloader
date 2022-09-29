// ----------------
// Fingerprints API @see https://docs.curseforge.com/?go#curseforge-core-api-fingerprints
// ----------------

package api

import (
	"bytes"
	"encoding/json"
	"github.com/lyssar/msdcli/curseforge/schemas"
	"net/http"
)

const (
	UriGetFingerprintsMatches      schemas.ApiUri = "/v1/fingerprints"
	UriGetFingerprintsFuzzyMatches schemas.ApiUri = "/v1/fingerprints/fuzzy"
)

// GetFingerprintsMatches @see https://docs.curseforge.com/?go#get-fingerprints-matches
func (api CurseforgeApi) GetFingerprintsMatches(request schemas.GetFingerprintMatchesRequestBody) (response schemas.GetFingerprintMatchesResponse, err error) {
	jsonRequest, err := json.Marshal(request)

	client := api.newCurseforgeClientForRoute(string(UriGetFingerprintsMatches))
	client.Method(http.MethodPost)
	client.Data(bytes.NewReader(jsonRequest))
	client.Request()
	err = client.GetContent(&response)

	return
}

// GetFingerprintsFuzzyMatches @see https://docs.curseforge.com/?go#get-fingerprints-fuzzy-matches
func (api CurseforgeApi) GetFingerprintsFuzzyMatches(request schemas.GetFuzzyMatchesRequestBody) (response schemas.GetFingerprintsFuzzyMatchesResponse, err error) {
	jsonRequest, err := json.Marshal(request)

	client := api.newCurseforgeClientForRoute(string(UriGetFingerprintsFuzzyMatches))
	client.Method(http.MethodPost)
	client.Data(bytes.NewReader(jsonRequest))
	client.Request()
	err = client.GetContent(&response)

	return
}
