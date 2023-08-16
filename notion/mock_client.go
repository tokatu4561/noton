package notion

import (
	"context"
	"net/http"
	"os"
)

type MockedClient struct {
	resFilePath string
}

func NewMockedClient (resFilePath string) *MockedClient {
	return &MockedClient{
		resFilePath: resFilePath,
	}
}

func (c *MockedClient) getPageId() string {
	return "mocked_page_id"
}

func (c *MockedClient) request(ctx context.Context, method, path string, body interface{}) (*http.Response, error) {
	b, err := os.Open(c.resFilePath)
	if err != nil {
		return nil, err
	}

	response := &http.Response{
		StatusCode: 200,
		Body:       b,
	}

	return response, nil
}