package utils

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func ApiCall(requestUrl string) []byte {
	apiClient := http.Client{
		Timeout: time.Second * 10, // Timeout after 10 seconds
	}

	req, err := http.NewRequest(http.MethodGet, requestUrl, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := apiClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				log.Fatal(err)
			}
		}(res.Body)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	return body
}
