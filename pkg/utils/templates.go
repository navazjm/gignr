package utils

import (
	"fmt"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

// $HOME/path/to/gignr/template.gitignore -> template.gitignore
func RemovePath(path string) string {
	filename := filepath.Base(path)
	return filename
}

// template.gitignore -> template
func RemoveGitignoreExt(filename string) string {
	return filename[:len(filename)-10]
}

// return [] of gitignore templates, where [i] = $HOME/path/to/gignr/templates/template.gitignore
func GetTemplates() []string {
	_, packagePath, _, ok := runtime.Caller(0) // packagePath = path/to/gignr/utils/template.go

	if !ok {
		fmt.Println("Error finding project directory")
	}

	templatesDir := path.Dir(packagePath)             // packagePath -> path/to/gignr/pkg/utils
	templatesDir = templatesDir[:len(templatesDir)-9] // remove utils from path -> path/to/gignr
	templatesDir += "templates/*.gitignore"           // set path to templates/*.gitignore -> path/to/gignr/templates/*.gitignore

	templateFiles, err := filepath.Glob(templatesDir) // get all .gitignore files in path/to/gignr/templates

	if err != nil {
		cobra.CheckErr(err)
	}

	return templateFiles
}

// $HOME/path/to/gignr/template.gitignore -> template
func ConvertPathsToFilenames(paths []string) []string {
	var filenames []string

	for _, filename := range paths {
		filename = RemovePath(filename)                          // remove $HOME/gignr/templates/ -> template.gitignore
		filename = RemoveGitignoreExt(filename)                  // remove file ext -> template
		filenames = append(filenames, strings.ToLower(filename)) // append lowercase template name
	}

	return filenames
}
