package more

import "io/ioutil"

var includeRegister = func(opts Opts) interface{} {
	return func(name string) string {
		path := getFilepath(opts, name)

		data, err := ioutil.ReadFile(path)
		panicError(err)

		return string(data)
	}
}
