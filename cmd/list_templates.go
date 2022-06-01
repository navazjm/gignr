package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
  "github.com/michaelnavs/gignr/utils"

)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all available templates to gnerate .gitignore file",
	Long:  "List all available templates to gnerate .gitignore file",
	Run:   onList,
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func onList(cmd *cobra.Command, args []string) {

  files := utils.GetTemplates() 


	for _, file := range files {
    fmt.Println(file)
	}
}
