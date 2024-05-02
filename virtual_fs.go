package skyclilib

import (
	"fmt"
	"os"
	"reflect"
	"text/template"
)

// Contains file template and it's values
type File struct {
	// if file is optional core will ask user if it's needed. Like "is {{ FileName }} needed? [Y/n]"
	//
	// if file is not optional core will ask only UserValues
	IsOptional bool

	// Nessary template values. They must be defined in plugin
	RequiredValues map[string]string

	// Values defined by user. Core will iterate over this map and ask user for it.
	// Example: "Enter {{ UserValue.Key }}: ".
	//
	// Keep values empty if plugin. If will be replaced
	UserValues map[string]string

	// File template
	// Template values as map[ValueName]Value. In template it's {{.ValueName}}
	Tpl string
}

// merge user and required values and return all this values
func (f *File) mergeValues() map[string]interface{} {
	allValues := make(map[string]interface{})
	for k, v := range f.UserValues {
		allValues[k] = v
	}
	for k, v := range f.RequiredValues {
		allValues[k] = v
	}
	return allValues
}

// JSON-like virtual project structure.
type FolderStructure map[string]interface{}

// JSON-like virtual project structure and IsOptional filed
// Must contain ONLY files and folders defined in this file. Any other files and folders will raise os.Exit(1)
type Folder struct {
	// if folder is optional core will ask user if it's needed. Like "is {{ FolderName }} needed? [Y/n]"
	// if user doesn't need this folder, any files and folders in this folder will be ignored
	IsOptional bool

	// Contains files and folders
	FolderStructure
}

// Generate project structure: create files and folders
//
// This func will recursively go around all defined files and folders. For folders it will create folder and call Gen().
// For files it will get user values with given func, render template and write it to file
func (f *Folder) Gen(pth string,
	askIfNeeded func(optionName string) bool,
	getUserFileConf func(filename string, userValues map[string]string) map[string]string,
) {
	for name, file := range f.FolderStructure {
		newPth := pth + "/" + name
		switch item := file.(type) {
		case File:
			if item.IsOptional && !askIfNeeded(name) {
				continue
			}
			item.UserValues = getUserFileConf(name, item.UserValues)
			renderTemplate(item.Tpl, newPth, item.mergeValues())
		case Folder:
			if item.IsOptional && !askIfNeeded(name) {
				continue
			}
			err := os.Mkdir(newPth, os.ModePerm)
			if err != nil {
				fmt.Println("can't create folder:", err)
				os.Exit(1)
			}
			item.Gen(newPth, askIfNeeded, getUserFileConf)
		default:
			fmt.Println("unknown type", reflect.TypeOf(item))
			fmt.Println("item:\n", item)
			os.Exit(1)
		}
	}
}

// Render template and write it to file
func renderTemplate(tpl, path string, data interface{}) {
	tplFile := mustGetFile(path)
	defer tplFile.Close()

	t := template.Must(template.New(path).Parse(tpl))
	err := t.Execute(tplFile, data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Create file and return it
// DOES NOT CLOSE FILE !
// If error occurs, os.Exit(1)
func mustGetFile(path string) *os.File {
	tplFile, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return tplFile
}
