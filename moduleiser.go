// The moduleiser adds a Modules directory to a .framework bundle, which exports everything
// in the header to the module.
package main

import (
	"flag"
	"os"
	"path/filepath"
	"text/template"
)

var (
	modmap = template.Must(template.New("modmap").Parse(`framework module {{.M}} {
  umbrella header "{{.H}}"

  export *
  module * { export * }
}
`))
	name   = flag.String("f", "demo.framework", "Name of the framework to add module map to")
	header = flag.String("h", "demo.h", "Name of the unbrella header file")
	module = flag.String("n", "demo", "Name of the module to add")
)

func main() {
	flag.Parse()
	md := filepath.Join(*name, "Versions/A/Modules")
	os.Mkdir(md, os.FileMode(0755))
	fn := filepath.Join(md, "module.modulemap")
	os.Remove(fn)
	f, err := os.Create(fn)
	if err != nil {
		panic(err)
	}
	x := struct{ H, M string }{
		H: *header,
		M: *module,
	}
	if err := modmap.Execute(f, x); err != nil {
		panic(err)
	}
	sl := filepath.Join(*name, "Modules")
	os.Remove(sl)
	if err := os.Symlink("Versions/A/Modules", sl); err != nil {
		panic(err)
	}
}
