# mule
A tool to be used with 'go generate' to embed external resources into Go code to create
single file exceutables without any dependencies.


## Scenario

An often used scenario in developing go applications is to embed external resources
to be able to create only one binary without any dependencies.
There are a number of existing packages solving this problem, like [bindata](https://github.com/a-urth/go-bindata),
[packr](https://github.com/gobuffalo/packr/tree/master/v2) or [packger](https://github.com/markbates/pkger)
and if you are looking for fancy features and unicorns you should probably better go there.
Usually they are creating a kind of virtual file system. Usually this realy alot more than I need for my
simple usecase on including one or two files into a small cli program.

This package 'mule' (the kinda donkey carrying huge loads) takes a _much simpler_ approach.
It just generates a single .go file for each resource you want to embed, including the
encoded resource wrapped in a function to access it.

It is intended to be run by go generate, though that is not required.


## Installation
   `go get github.com/wlbr/mule`


## Usage

Simply add a line

   `//go:generate mule mybinary.file`

to one of your source file for each resource you want to embed. Every time you run a 'go generate' in the
corresponding folder, the file 'mybinary.go' will be created. It contains a
function 'mybinaryResource' returning the resource as a []byte.

See [mulex.go](https://github.com/wlbr/mule/blob/master/example/mulex.go) for a very, very simple example.

You may use 'mule mytbinary.file' directly on the command line.


## Switches

Usage of mule: `mule [switches] resourcefilename`<br>
   -e<br>
      export the generated, the resource returning function. Default (false) means
      the function will not be exported.

   -f<br>
      no formatting of the generated source. Default false means source will be
      formatted with gofmt.

   -n string<br>
    	 name of generated, the resource returning function. Its name will have
      'Resource' attached. Will be set to $(basename -s .ext outputfile) if empty
      (default). Take care of "-" within the name, especially when the name is
      calculated from the resources file name.  A '-' would create an invalid go
      function name

   -o string<br>
    	 name of output file. Defaults to name of resource file excluding
      extension + '.go'.

   -p string<br>
  	 name of package to be used in generated code (default "main").

   -t string<br>
    	 name of alternate code generation template file. If empty (default), then
      the embedded template will be used. Template variables supplied are:
      .Name, .Package, .Content


## Code
* Documentation: https://godoc.org/github.com/wlbr/mule
* Lint: http://go-lint.appspot.com/github.com/wlbr/mule
* Continous Integration: [![Travis Status](https://api.travis-ci.com/wlbr/mule.svg?branch=master)](https://travis-ci.com/wlbr/mule)
* Test Coverage: [![Coverage Status](https://coveralls.io/repos/github/wlbr/mule/badge.svg?branch=master)](https://coveralls.io/github/wlbr/mule?branch=master)
* Metrics: [![GoReportCard](https://goreportcard.com/badge/github.com/wlbr/mule)](https://goreportcard.com/report/github.com/wlbr/mule)
