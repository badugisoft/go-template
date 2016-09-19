package funcs

func init() {
	registerFunc("gtMakeMap",
		func(args ...interface{}) map[string]interface{} {
			ret := map[string]interface{}{}
			for i, end := 0, len(args); i < end; i += 2 {
				ret[args[i].(string)] = args[i+1]
			}
			return ret
		})
}
