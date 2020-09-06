package messages

import (
	"bytes"
	"io/ioutil"
	"log"
	"path/filepath"
	"runtime"
	"text/template"
)

//
type Template struct {
	Name     string
	T        *template.Template
	FuncMap  *template.FuncMap
	Logger   bool
	Debugger bool
}

func (Template *Template) GetT() *template.Template{
	return Template.T
}

func (Template *Template) SetLogger(val bool) *Template{
	Template.Logger = val
	return Template
}

func (Template *Template) SetDebugger(val bool) *Template{
	Template.Debugger = val
	return Template
}

// New messages template
func New(Name string) *Template {
	return &Template{Name: Name}
}

func ( Template *Template)SetFuncMap(funcMap template.FuncMap)*Template{
	Template.FuncMap = &funcMap
	return Template
}

// ReLoad templates form folder
// path - templates files path
// ext	- template file extension
func (Template *Template) ReLoad(path string, ext string) (err error) {
	filesInfo, err := ioutil.ReadDir(path)
	if err != nil {
		return
	}
	filesPath := []string{}
	for _, fileInfo := range filesInfo {
		if !fileInfo.IsDir() {
			Ext:=filepath.Ext(fileInfo.Name())
			if Ext == ext {
				filesPath = append(filesPath, path+"/"+fileInfo.Name())
			}
		}
	}
	Template.T, err = template.
		New(Template.Name).
		ParseFiles(filesPath...)
	return
}

//Execute Template
func (Template Template) Execute(name string, source interface{}) (result string) {
	var res []byte
	buf := bytes.NewBuffer(res)
	err := Template.GetT().ExecuteTemplate(buf, name, source)
	if err != nil {
		return err.Error()
	}
	return buf.String()
}

func (Template Template) Debug(name string, source interface{}){
	_, filename, line, _ := runtime.Caller(1)
	if Template.Debugger{
		log.Println(Template.Execute(name, struct{
			Source interface{}
			Filename  string
			Line int
		}{source, filename, line}))
	}
}

func (Template Template) Log(name string, source interface{}){
	if Template.Logger{
		log.Println(Template.Execute(name, source))
	}
}
