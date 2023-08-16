package notion_test

import (
	"testing"

	"github.com/tokatu4561/noton/notion"
)

func TestBlockClient(t *testing.T) {

	t.Run("AppendBlock", func(t *testing.T) {
		tests := []struct {
			name       	string
			request    	*notion.AppendBlockParams
			resFilePath string
			want       	*notion.AppendBlockResponse
			wantErr    	bool
			err        	error
		}{
			{
				name:       "append paragraph text in page",
				resFilePath: "test_data/append_block.json",
				request: &notion.AppendBlockParams{
					Children: []interface{}{
						notion.ParagraphBlock{
							Block: notion.Block{
								Object: "block",
								Type: "paragraph",
							},
							Paragraph: notion.Paragraph{
								RichText: []notion.RichText{
									{
										Text: &notion.Text{
											Content: "test",
										},
									},
								},
							},
						},
					},
				},
				want: &notion.AppendBlockResponse{
					Object: "block",
					Results: []interface{}{
						&notion.ParagraphBlock{
							Block: notion.Block{
								Object: "block",
								Type: "paragraph",
							},
							Paragraph: notion.Paragraph{
								RichText: []notion.RichText{
									{
										Text: &notion.Text{
											Content: "test",
										},
									},
								},
							},
						},
					},
				},
				wantErr: false,
				err:     nil,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				client := notion.NewMockedClient(tt.resFilePath)
				pageClient := notion.NewPageClient(client)
				got, err := pageClient.AppendBlock(*tt.request)

				if (err != nil) != tt.wantErr {
					t.Errorf("AppendBlock() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

				if got.Object != tt.want.Object {
					t.Errorf("AppendBlock() got = %v, want %v", got, tt.want)
				}
			})
		}
	})
}