package utils

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func GetTemplates() []string {
	path, err := os.Getwd()

	if err != nil {
		cobra.CheckErr(err)
	}

	path += "/templates/"

	files, err := filepath.Glob(path + "*.gitignore")

	if err != nil {
		cobra.CheckErr(err)
	}

	for i, file := range files {
		filename := filepath.Base(file)
		files[i] = filename[:len(filename)-10] // removes .gitignore from file name
	}

	return files
}
