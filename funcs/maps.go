package funcs

import "text/template"

var _map = template.FuncMap{}

func register(name string, f interface{}) {
	_map[name] = f
}

func GetMap() template.FuncMap {
	return _map
}
