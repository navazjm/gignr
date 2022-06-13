package commands

import (
	"bufio"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/navazjm/gignr/pkg/utils"
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
	templatePaths := utils.GetTemplates()                             // used to read content of gitignore templates
	templateFilenames := utils.ConvertPathsToFilenames(templatePaths) // used to match specified templates by user

	var gitignoreContents []string // contents of the new gitignore file

	for _, template := range templates {
		template = strings.ToLower(template)
		for i, file := range templateFilenames {
			if template != file {
				continue
			}

			templateFile, err := os.Open(templatePaths[i])

			if err != nil {
				cobra.CheckErr(err)
			}

			defer templateFile.Close()

			scanner := bufio.NewScanner(templateFile)
			for scanner.Scan() {
				line := scanner.Text()
				line += "\n"
				gitignoreContents = append(gitignoreContents, line)
			}

			if err := scanner.Err(); err != nil {
				cobra.CheckErr(err)
			}
		}
	}

	// get gitignore path in users current working dir
	gitignorePath, err := os.Getwd()
	if err != nil {
		cobra.CheckErr(err)
	}
	gitignorePath += "/.gitignore"

	var gitignoreFile *os.File

	if isAppending {
		gitignoreFile, err = os.OpenFile(gitignorePath, os.O_APPEND|os.O_WRONLY, 0644)
	} else {
		gitignoreFile, err = os.Create(gitignorePath)
	}

	if err != nil {
		cobra.CheckErr(err)
	}

	defer gitignoreFile.Close()

	for _, line := range gitignoreContents {
		if _, err := gitignoreFile.WriteString(line); err != nil {
			cobra.CheckErr(err)
		}
	}
}
