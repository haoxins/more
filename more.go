package more

import m2s "github.com/mitchellh/mapstructure"
import . "github.com/tj/go-debug"
import "path/filepath"
import "text/template"
import "io/ioutil"
import "bytes"
import "os"

var debug = Debug("more")

func New(options map[string]interface{}) *Render {
	var opts Opts

	err := m2s.Decode(options, &opts)
	panicError(err)

	debug("opts: %v", opts)

	if opts.Ext == "" {
		opts.Ext = ".html"
	}

	wd, err := os.Getwd()
	panicError(err)

	if opts.Dir == "" {
		opts.Dir = wd
	} else {
		opts.Dir = filepath.Join(wd, opts.Dir)
	}

	r := Render{
		opts:  opts,
		funcs: template.FuncMap{},
		cache: map[string]*template.Template{},
	}

	r.init()

	return &r
}

type Opts struct {
	Dir   string
	Ext   string
	Cache bool
}

type Render struct {
	opts  Opts
	funcs template.FuncMap
	cache map[string]*template.Template
}

func (r Render) init() {
	r.Register("include", includeRegister)
}

func (r Render) load(name string) *template.Template {
	tmpl, ok := r.cache[name]
	if ok && r.opts.Cache {
		return tmpl
	}

	path := getFilepath(r.opts, name)

	data, err := ioutil.ReadFile(path)
	panicError(err)

	tmpl, err = template.New(name).Funcs(r.funcs).Parse(string(data))
	panicError(err)

	r.cache[name] = tmpl

	return tmpl
}

func (r Render) Render(name string, data interface{}) string {
	tmpl := r.load(name)

	var buf bytes.Buffer
	tmpl.Execute(&buf, data)

	return buf.String()
}

func (r Render) Register(name string, reg register) {
	r.funcs[name] = reg(r.opts)
}

type register func(opts Opts) interface{}
