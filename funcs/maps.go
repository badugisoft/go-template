package funcs

import "text/template"

var FuncMap = template.FuncMap{
	"gtMakeMap":   MakeMap,
	"gtMakeSlice": MakeSlice,
}
