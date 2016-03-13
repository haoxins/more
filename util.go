package more

import "path/filepath"

func panicError(err error) {
	if err != nil {
		panic(err)
	}
}

func getFilepath(opts Opts, name string) string {
	path := filepath.Join(opts.Dir, name)

	if filepath.Ext(path) == "" {
		path = path + "." + opts.Ext
	}

	return path
}
