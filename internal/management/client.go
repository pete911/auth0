package management

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	apiVersion  = "v2"
	httpTimeout = 5 * time.Second
)

type Client struct {
	url        string
	token      string
	httpClient *http.Client
}

// NewClient returns auth0 client for supplied domain and management api access token
func NewClient(domain, clientId, clientSecret string) (Client, error) {
	domain = strings.Trim(strings.TrimPrefix(domain, "https://"), "/")
	httpClient := &http.Client{Timeout: httpTimeout}
	token, err := getToken(httpClient, domain, clientId, clientSecret)
	if err != nil {
		return Client{}, err
	}

	url := fmt.Sprintf("https://%s/api/%s", domain, apiVersion)
	return Client{
		url:        url,
		token:      token.AccessToken,
		httpClient: &http.Client{Timeout: httpTimeout},
	}, nil
}

// Do send request to auth0, path is http path without '<domain>/api/<version/' prefix. User can provide optional
// url query parameters, json request and expected json response. Returned error can be type of ErrHttp if the
// returned response status code is not 2xx
func (c Client) Do(method, path string, queryParams map[string]string, request, response interface{}) error {
	requestUrl := fmt.Sprintf("%s/%s", c.url, strings.TrimPrefix(path, "/"))
	req, err := getRequest(c.token, method, requestUrl, queryParams, request)
	if err != nil {
		return fmt.Errorf("get reqeust: %w", err)
	}
	return sendRequest(c.httpClient, req, response)
}

func getToken(httpClient *http.Client, domain, clientId, clientSecret string) (TokenResponse, error) {
	requestBody := NewTokenRequest(domain, clientId, clientSecret)
	requestUrl := fmt.Sprintf("https://%s/oauth/token", domain)
	request, err := getRequest("", http.MethodPost, requestUrl, nil, requestBody)
	if err != nil {
		return TokenResponse{}, fmt.Errorf("request: %w", err)
	}

	var response TokenResponse
	err = sendRequest(httpClient, request, &response)
	return response, err
}

func getRequest(token, method, requestUrl string, queryParams map[string]string, request interface{}) (*http.Request, error) {

	if len(queryParams) != 0 {
		u, err := url.Parse(requestUrl)
		if err != nil {
			return nil, fmt.Errorf("request url %s: %w", requestUrl, err)
		}
		values := url.Values{}
		for k, v := range queryParams {
			values.Add(k, v)
		}
		u.RawQuery = values.Encode()
		requestUrl = u.String()
	}

	var requestBody io.Reader
	if request != nil {
		b, err := json.Marshal(request)
		if err != nil {
			return nil, fmt.Errorf("marshal request body: %w", err)
		}
		requestBody = bytes.NewBuffer(b)
	}

	req, err := http.NewRequest(method, requestUrl, requestBody)
	if err != nil {
		return nil, err
	}

	if token != "" {
		req.Header.Set("authorization", fmt.Sprintf("Bearer %s", token))
	}
	req.Header.Set("content-type", "application/json")
	return req, nil
}

func sendRequest(httpClient *http.Client, request *http.Request, response interface{}) error {
	resp, err := httpClient.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("io read body: %w", err)
	}

	if resp.StatusCode/100 != 2 {
		return NewErrHttp(request.Method, request.URL.String(), resp.StatusCode, body)
	}

	if response != nil {
		if err := json.Unmarshal(body, response); err != nil {
			return fmt.Errorf("unmarshal %s response: %w", request.URL, err)
		}
	}
	return nil
}
