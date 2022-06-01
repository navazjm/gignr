package utils

import (
	"fmt"
	"path"
	"path/filepath"
	"runtime"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

// returns an array of all .gitignore templates found in path/to/gignr/templates/
func GetTemplates() []string {
	_, packagePath, _, ok := runtime.Caller(0) // packagePath -> path/to/gignr/utils/

	if !ok {
		fmt.Println("Error finding project directory")
	}

	templatesPath := path.Dir(packagePath)

	templatesPath = templatesPath[:len(templatesPath)-5] // remove utils from path -> path/to/gignr
	templatesPath += "templates/*.gitignore"             // set path to templates/*.gitignore -> path/to/gignr/templates/*.gitignore

	files, err := filepath.Glob(templatesPath) // get all .gitignore files in path/to/gignr/templates

	if err != nil {
		cobra.CheckErr(err)
	}

	for i, file := range files {
		filename := filepath.Base(file)        // remove $HOME/gignr/templates/ -> template.gitignore
		filename = filename[:len(filename)-10] // removes .gitignore from file name -> template
		files[i] = strings.ToLower(filename)   // converts to lowercase
	}

	sort.Strings(files) // alphabetically sort the templates

	return files
}
