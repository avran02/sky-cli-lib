package skyclilib

import (
	"fmt"
	"os"
	"reflect"
	"text/template"
)

type PluginConfiger interface {
	GetAvailalableOptions() []string
	GetVirtualFs(neededOptions map[string]bool) *Folder
	GetRequiredParams() []string
}

type File struct {
	tpl string
}

type Folder map[string]interface{}

func (f *Folder) Gen(pth string, conf PluginConfiger) {
	for name, file := range *f {
		newPth := pth + "/" + name
		switch item := file.(type) {
		case File:
			RenderTemplate(item.tpl, newPth, conf)
		case Folder:
			err := os.Mkdir(newPth, os.ModePerm)
			if err != nil {
				fmt.Println("can't create folder:", err)
				os.Exit(1)
			}
			item.Gen(newPth, conf)
		default:
			fmt.Println("unknown type", reflect.TypeOf(item))
			os.Exit(1)
		}
	}
}

func RenderTemplate(tpl, path string, data interface{}) {
	tplFile := mustGetFile(path)
	defer tplFile.Close()

	t := template.Must(template.New(path).Parse(tpl))
	err := t.Execute(tplFile, data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func mustGetFile(path string) *os.File {
	tplFile, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return tplFile
}
