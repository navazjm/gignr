package utils

import (
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

func GetTemplates() []string {
	// TODO: change os.GetWD(), this only works if the user is inside project's root dir.
	path, err := os.Getwd() // set path to location of gignr

	if err != nil {
		cobra.CheckErr(err)
	}

	path += "/templates/*.gitignore"

	files, err := filepath.Glob(path)

	if err != nil {
		cobra.CheckErr(err)
	}

	for i, file := range files {
		filename := filepath.Base(file)        // remove $HOME/gignr/templates/ -> template.gitignore
		filename = filename[:len(filename)-10] // removes .gitignore from file name -> template
		files[i] = strings.ToLower(filename)
	}

	sort.Strings(files)

	return files
}
