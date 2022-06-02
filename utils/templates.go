package utils

import (
	"fmt"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

func RemovePath(path string) string {
	filename := filepath.Base(path)
	return filename
}

func RemoveGitignoreExt(filename string) string {
	return filename[:len(filename)-10] // removes .gitignore from file name -> template
}

// returns an array of all .gitignore templates found in path/to/gignr/templates/
func GetTemplates() []string {
	_, packagePath, _, ok := runtime.Caller(0) // packagePath -> path/to/gignr/utils/template.go

	if !ok {
		fmt.Println("Error finding project directory")
	}

	templatesDir := path.Dir(packagePath)             // packagePath -> path/to/gignr/utils
	templatesDir = templatesDir[:len(templatesDir)-5] // remove utils from path -> path/to/gignr
	templatesDir += "templates/*.gitignore"           // set path to templates/*.gitignore -> path/to/gignr/templates/*.gitignore

	templateFiles, err := filepath.Glob(templatesDir) // get all .gitignore files in path/to/gignr/templates

	if err != nil {
		cobra.CheckErr(err)
	}

	return templateFiles
}

func ConvertPathsToFilenames(paths []string) []string {
	var filenames []string

	for _, filename := range paths {
		filename = RemovePath(filename)                          // remove $HOME/gignr/templates/ -> template.gitignore
		filename = RemoveGitignoreExt(filename)                  // remove file ext -> template
		filenames = append(filenames, strings.ToLower(filename)) // append lowercase template name
	}

	return filenames
}
