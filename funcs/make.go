package funcs

func init() {
	register("gtMakeMap", MakeMap)     // legacy
	register("gtMakeSlice", MakeSlice) // legacy

	register("makeMap", MakeMap)
	register("makeSlice", MakeSlice)
}

func MakeMap(args ...interface{}) map[string]interface{} {
	ret := map[string]interface{}{}
	for i, end := 0, len(args); i < end; i += 2 {
		ret[args[i].(string)] = args[i+1]
	}
	return ret
}

func MakeSlice(args ...interface{}) []interface{} {
	return args
}
