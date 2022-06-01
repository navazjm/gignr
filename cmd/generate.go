package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var templates []string
var isAppending bool

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a new .gitignore file.",
	Long:  "Generate a new .gitignore file. This will overwrite an existing .gitignore file by default. See append flag for modifying existing .gitignore files.",
	Run:   onGnerate,
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	generateCmd.Flags().BoolVarP(&isAppending, "append", "a", false, "This will append templates to an existing .gitignore file.")
	generateCmd.Flags().StringArrayVarP(&templates, "template", "t", templates, "REQUIRED. Specify which templates to use to generate .gitignore file.")
	generateCmd.MarkFlagRequired("template")
}

func onGnerate(cmd *cobra.Command, args []string) {
	// TODO: generate .gitignore file based on passed in templates
	if isAppending {
		fmt.Println("we are appending")
	}
	fmt.Println(templates)
}
