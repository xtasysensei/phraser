package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

func init() {
	rootCmd.AddCommand(gendocCmd)
}

var gendocCmd = &cobra.Command{
	Use:   "gendoc",
	Short: "Generate Markdown documentation for all commands",
	Run: func(cmd *cobra.Command, args []string) {
		// Ensure the 'docs' directory exists
		if err := os.MkdirAll("./docs", os.ModePerm); err != nil {
			cobra.CheckErr(err)
		}

		// Generate the Markdown documentation
		err := doc.GenMarkdownTree(rootCmd, "./docs")
		if err != nil {
			cobra.CheckErr(err)
		}
		fmt.Println("Documentation generated in the 'docs' directory.")
	},
}
