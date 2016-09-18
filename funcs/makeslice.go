package funcs

func init() {
	registerFunc("gtMakeSlice",
		func(args ...interface{}) []interface{} {
			return args
		}
	)
}
