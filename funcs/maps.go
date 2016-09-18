package funcs

import "text/template"

var FuncMap = template.FuncMap{}

func registerFunc(name string, f interface{}) {
	FuncMap[name] = f
}
