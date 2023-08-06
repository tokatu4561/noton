package notion

import (
	"net/http"
	"net/url"
)

const (
	apiUrl = "https://api.notion.com/v1"
	apiKey = "test"
	pageId = "test"
	version = "2022-06-28"
)

type Client struct {
	httpClient  *http.Client
	baseUrl     *url.URL
	version  	string
	token 		string
	Page        string
}

func NewClient() *Client {
	client := &http.Client{}

	url, err := url.Parse(apiUrl)
	if err != nil {
		panic(err)
	}

	c := &Client{
		httpClient:  client,
		token:       apiKey,
		baseUrl:     url,
		version: 	 version,
	}

	return c
}

func (c *Client) request(method, path string, body interface{}) (*http.Response, error) {
	requestUrl := c.baseUrl.String() + path
	req, err := http.NewRequest(method, requestUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer" + c.token)
	req.Header.Add("Notion-Version", c.version)
	req.Header.Add("Content-Type", "application/json")

	response, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	
	if response.StatusCode != http.StatusOK {
		return nil, err
	}

	return response, nil
}