// Mule is a tool to be used with 'go generate' to embed external resources files
// into Go code and therefore into the resulting executable.
//
//
// Scenario
//
// An often used scenario in developing go applications is to embed external resources
// to be able to create only one binary without any dependencies.
// There are a number of existing packages solving this problem, like bindata (https://github.com/a-urth/go-bindata),
// packr (https://github.com/gobuffalo/packr/tree/master/v2) or packger (https://github.com/markbates/pkger)
// and if you are looking for fancy features and unicorns you should probably better go there.
// Usually they are creating a kind of virtual file system.
// Generelly this really a lot more than I need for my
// simple usecase in including one or two files into a small cli program.
//
// Compared to that mule is extremely simple. The only thing you need to embed a file
// to your code is one line in your code (a go generate command).
// And you need just another one line to access the embedded file from your code.
//
// See https://github.com/wlbr/mule/blob/master/example/mulex.go for a very, very
// simple example.
//
// It is intended to be run by go generate.
//
//
// Usage
//
// Simply add a line
//    //go:generate mule mybinary.file
//
// for each resource you want to embed. Every time you run a 'go generate' in the
// corresponding folder, the file 'mybinary.go' will be created. It contains a
// function 'mybinaryResource' returning the resource as a []byte.
//
// You may use 'mule mytbinary.file' directly on the command line.
//
//
// Switches
//
// Usage of mule: 'mule [switches] resourcefilename'
//
//  -e
//     export the generated, the resource returning function. Default (false) means
//     the function will not be exported.
//
//  -f
//     no formatting of the generated source. Default false means source will be
//     formatted with gofmt.
//
//  -n string
//     name of generated, the resource returning function. Its name will have
//     'Resource' attached. Will be set to $(basename -s .ext outputfile) if empty
//     (default). Take care of "-" within the name, especially when the name is
//     calculated from the resources file name.  A '-' would create an invalid go
//     function name
//
//  -o string
//     name of output file. Defaults to name of resource file excluding
//     extension + '.go'.
//
//  -p string
//     name of package to be used in generated code (default "main").
//
//  -t string
//     name of alternate code generation template file. If empty (default), then
//     the embedded template will be used. Template variables supplied are:
//     .Name, .Package, .Content
//
package main

//go:generate templify -p main -o embed.go embed.tpl

import (
	"encoding/base64"
	"errors"
	"flag"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"text/template"
)

var pckg string
var inputfile string
var outfilename string
var functionname string
var tmplname string
var frmt bool
var exp bool

func flagging() {
	flag.StringVar(&pckg, "p", "main", "name of package to be used in generated code.")
	flag.StringVar(&outfilename, "o", "", "name of output file. Defaults to name of resource file + '.go'.")
	flag.StringVar(&functionname, "f", "", "name of generated, the resource returning function. Its name will "+
		"have 'Resource' attached. Will be set to $(basename -s .ext outputfile) if empty (default).")
	flag.StringVar(&tmplname, "t", "", "name of alternate code generation template file. If empty (default), "+
		"then the embedded template will be used. Template variables supplied are: .Name, .Package, .Content")
	flag.BoolVar(&frmt, "n", false, "do not format the generated source. Default false means source will be formatted.")
	flag.BoolVar(&exp, "e", false, "export the generated, the resource returning function. "+
		"Default (false) means the function will not be exported.")
	flag.Parse()

	inputfile = flag.Arg(0)
	if inputfile == "" {
		fmt.Println(errors.New("no resource file given as argument"))
		os.Exit(1)
	}

	if outfilename == "" {
		indir := path.Dir(inputfile)
		inext := path.Ext(path.Base(inputfile))
		inname := strings.TrimSuffix(path.Base(inputfile), inext)
		outfilename = fmt.Sprintf("%s/%s.go", indir, inname)
	}

	if functionname == "" {
		ext := path.Ext(path.Base(outfilename))
		functionname = strings.TrimSuffix(path.Base(outfilename), ext)
	}

	if exp {
		functionname = strings.ToUpper(functionname[0:1]) + functionname[1:]
	}
}

func readMuleTemplate(tplname string) (*template.Template, error) {
	tpl, err := template.ParseFiles(tplname)
	if err != nil {
		fmt.Printf("Error reading muletemplate file '%s'\n%v", tplname, err)
	}
	return tpl, err
}

func readTargetResource(resname string) string {
	res, err := ioutil.ReadFile(resname)
	encoded := base64.StdEncoding.EncodeToString(res)
	if err != nil {
		fmt.Printf("Error reading target resource file '%s'\n%v", resname, err)
		os.Exit(1)
	}

	return encoded
}

func formatFile(fname string) {
	fstr, err := ioutil.ReadFile(fname)
	if err != nil {
		fmt.Printf("Error reading generated file %s before passing it to gofmt.\n%v\n", fname, err)
		os.Exit(1)
	} else {
		fstr, err = format.Source(fstr)
		if err != nil {
			fmt.Printf("Error running gofmt on the generated file '%s'\n%v\n", fname, err)
			os.Exit(1)
		} else {
			foutfile, err := os.Create(fname)
			if err != nil {
				fmt.Printf("Error creating formatted target file '%s'\n%v\n", fname, err)
				os.Exit(1)
			} else {
				defer foutfile.Close()
				fmt.Fprintf(foutfile, "%s", fstr)
			}
		}
	}
}

func main() {
	flagging()

	var tpl *template.Template
	var err error

	if tmplname != "" {
		tpl, err = readMuleTemplate(tmplname)

	} else {
		tpl, err = template.New("embed").Parse(embedTemplate())
	}

	if err != nil {
		fmt.Printf("Error parsing code generation template\n%v", err)
		os.Exit(1)
	}

	data := struct {
		Package string
		Name    string
		Content string
	}{
		Package: pckg,
		Name:    strings.Split(functionname, ".")[0],
	}
	data.Content = readTargetResource(inputfile)

	outfile, err := os.Create(outfilename)
	if err != nil {
		fmt.Printf("Error creating target file '%s'\n%v\n", outfilename, err)
		os.Exit(1)
	}
	defer outfile.Close()
	tpl.Execute(outfile, data)

	if !frmt {
		formatFile(outfilename)
	}

}
