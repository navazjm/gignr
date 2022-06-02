package cmd

import (
	"fmt"
	"sort"

	"github.com/michaelnavs/gignr/utils"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all available templates to generate .gitignore file",
	Long:  "List all available templates to generate .gitignore file",
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

	templateFilenames := utils.GetTemplates()
  templateFilenames = utils.ConvertPathsToFilenames(templateFilenames)

  sort.Strings(templateFilenames)

  for _, templateFilename := range templateFilenames {
    fmt.Println(templateFilename)
  }
}
