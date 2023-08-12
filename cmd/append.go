/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tokatu4561/noton/notion"
)

// appendCmd represents the append command
var appendCmd = &cobra.Command{
	Use:   "append",
	Short: "append a block to a page",
	Long: `append a block to a page`,
	Run: func(cmd *cobra.Command, args []string) {
		text, _ := cmd.Flags().GetString("text")

		notionParagraph := notion.ParagraphBlock{
			Block: notion.Block{
				Object: "block",
				Type: "paragraph",
			},
			Paragraph: notion.Paragraph{
				RichText: []notion.RichText{
					{
						Text: &notion.Text{
							Content: text,
						},
					},
				},
			},
		}

		params := notion.AppendBlockParams{
			Children: []interface{}{
				notionParagraph,
			},
		}

		result, err := notion.AppendBlock(params)
		if err != nil {
			panic(err)
		}
		
		fmt.Println("append called with text: %s result: %s" + text, result)
	},
}

func init() {
	rootCmd.AddCommand(appendCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	appendCmd.PersistentFlags().StringP("text", "t", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// appendCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
