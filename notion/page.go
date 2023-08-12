package notion

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type BlockID string

type BlockInterface interface {
	GetType() Block
	GetCreatedTime() *time.Time
}

// See https://developers.notion.com/reference/block
type Block struct {
	// ID             BlockID    `json:"id,omitempty"`
	Object         string `json:"object"`
	Type           string  	  `json:"type"`
	CreatedTime    *time.Time `json:"created_time,omitempty"`
}

func (b Block) GetType() string {
	return b.Type
}

func (b Block) GetCreatedTime() *time.Time {
	return b.CreatedTime
}

type Text struct {
	Content string `json:"content"`
}

type RichText struct {
	Text        *Text   `json:"text,omitempty"`
	PlainText   string  `json:"plain_text,omitempty"`
}

type Paragraph struct {
	RichText []RichText `json:"rich_text"`
	Children []Block    `json:"children,omitempty"`
	Color    string     `json:"color,omitempty"`
}

type ParagraphBlock struct {
	Block
	Paragraph Paragraph `json:"paragraph"`
}

type AppendBlockParams struct {
	// Child content to append to a container block as an array of block objects.
	Children []any `json:"children"`
}

func AppendBlock(params AppendBlockParams) (result *AppendBlockResponse, err error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	
	client := NewClient()
	requestPath := fmt.Sprintf("/blocks/%s/children", client.page)
	requestBody := params
	
	res, err := client.request(ctx, http.MethodPatch, requestPath, requestBody)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		var errorResponse ErrorResponse
		err = json.NewDecoder(res.Body).Decode(&errorResponse)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("API response is not OK: %s Message: %s", res.Status, errorResponse.Message)
	}

	var responseResults AppendBlockResponse
	err = json.NewDecoder(res.Body).Decode(&responseResults)
	if err != nil {
		return nil, err
	}

	return &responseResults, nil
}

type AppendBlockResponse struct {
	Object  string `json:"object"`
	Results []Block    `json:"results"`
}

type ErrorResponse struct {
	Object string `json:"object"`
	Status int    `json:"status"`
	Code string `json:"code"`
	Message string `json:"message"`
}