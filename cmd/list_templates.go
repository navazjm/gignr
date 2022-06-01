package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
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
	path, err := os.Getwd()

	if err != nil {
		cobra.CheckErr(err)
	}

	path += "/templates/"

	files, err := filepath.Glob(path + "*.gitignore")

	if err != nil {
		cobra.CheckErr(err)
	}


	for _, file := range files {
    filename := filepath.Base(file)
    filename = filename[: len(filename)-10] // removes .gitignore from file name
    fmt.Println(filename)
	}
}
