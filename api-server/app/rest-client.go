package app

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	AUTH_HEADER     = "Authorization"
	AUTH_BASIC_TYPE = "Basic "
	CONTENT_TYPE    = "Content-Type"
)

type HttpClient struct {
	Client *http.Client
}

func (h *HttpClient) HEAD(url string) (int, error) {
	r, err := http.NewRequest(http.MethodHead, url, nil)
	//r.Header.Set(AUTH_HEADER, AUTH_BASIC_TYPE+c.B64)

	res, err := h.Client.Do(r)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return res.StatusCode, nil
}

func (h *HttpClient) GET(url string) (string, int) {
	r, err := http.NewRequest(http.MethodGet, url, nil)

	res, err := h.Client.Do(r)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return "", http.StatusInternalServerError
	}
	defer res.Body.Close()
	if res.StatusCode == http.StatusOK || res.StatusCode == http.StatusCreated {
		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return "", http.StatusInternalServerError
		}
		body := string(data)
		return body, http.StatusOK
	}

	switch res.StatusCode {
	case http.StatusForbidden:
		fmt.Printf("Use do not have the right access to this resource")
		return "", http.StatusForbidden
	case http.StatusUnauthorized:
		fmt.Printf("User is not authenticated or has wrong credentials")
		return "", http.StatusUnauthorized
	default:
		fmt.Printf("Internal Server Error occured")
		return "", http.StatusInternalServerError
	}
}
