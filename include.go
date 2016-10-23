package more

import "io/ioutil"

var includeRegister = func(opts Opts) interface{} {
	var cache []byte

	return func(name string) string {
		path := getFilepath(opts, name)

		if opts.Cache {
			if len(cache) == 0 {
				var e error
				cache, e = ioutil.ReadFile(path)
				panicError(e)
			}
			return string(cache)
		}

		data, err := ioutil.ReadFile(path)
		panicError(err)

		return string(data)
	}
}
