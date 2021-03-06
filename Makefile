
all: generate test build		## a simple call of 'make' without arguments will do everything.

.PHONY: clean
clean:		## tidy up everything.
	rm -f mule
	#rm -f embed.go
	rm -f example/mulex
	rm -f example/gopher.go
	rm -f example/example
	rm -f example/gopherexported.jpg



generate: embed.tpl		## generate the code generation template (called by build).
	go generate mule.go

build: generate mule.go embed.go		## create executable.
	go build mule.go embed.go

test: generate mule.go embed.tpl embed.go mule_test.go embed_test.go		## run the tests.
	go test

examples: example/mulex.go example/gopher.jpg example/gopher.go  	 	## compile the example in its subfolder.
	go generate example/mulex.go
	go build example/mulex.go example/gopher.go

dep:		## install all needed dependencies
	go get github.com/wlbr/templify



help: ## this help message
	@echo "$$(grep -hE '^\S+:.*##' $(MAKEFILE_LIST) | sed -e 's/:.*##\s*/:/' -e 's/^\(.\+\):\(.*\)/\\x1b[36m\1\\x1b[m:\2/' | column -c2 -t -s :)"


