package notion

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"

	"context"
)

const (
	apiUrl = "https://api.notion.com/v1"
	version = "2022-06-28"
)

type ClientInterface interface {
	getPageId() string
	request(ctx context.Context, method, path string, body interface{}) (*http.Response, error)
}

type Client struct {
	httpClient  *http.Client
	baseUrl     *url.URL
	version  	string
	token 		string
	page        string
}

// NewClient returns a new Notion API client.
func NewClient() *Client {
	client := &http.Client{}

	url, err := url.Parse(apiUrl)
	if err != nil {
		panic(err)
	}

	config, err := loadConfig()
	if err != nil {
		panic(err)
	}

	c := &Client{
		httpClient:  client,
		token:       config.secretToken,
		baseUrl:     url,
		page: 	  	 config.pageId,
		version: 	 version,
	}

	return c
}

// getPageId returns the page id.
func (c *Client) getPageId() string {
	return c.page
}

// request sends a request to the Notion API and returns the response.
func (c *Client) request(ctx context.Context, method, path string, body interface{}) (*http.Response, error) {
	requestUrl := c.baseUrl.String() + path

	bodyJson, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	
	req, err := http.NewRequest(method, requestUrl, bytes.NewBuffer(bodyJson))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer " + c.token)
	req.Header.Add("Notion-Version", c.version)
	req.Header.Add("Content-Type", "application/json")

	response, err := c.httpClient.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}

	return response, nil
}