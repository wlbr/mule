package main

//go:generate mule gopher.jpg
// the line above executes mule to generate a go source code representation
// of gopher.jpg, whenever you run `go generate` in this directory.

import (
	"fmt"
	"os"
)

func main() {
	// once you did a 'go generate' and then a 'go build' in this folder, you will have
	// an executable 'example' that has the image gopher.jpg included. (See line 3 above).
	//
	// With the follwing code we access this embedded resource and write it to disk.

	gopher, err := gopherResource() //decoding the embedded resource

	if err != nil {
		fmt.Printf("Error decoding the embedded resource.\n%v\n", err)
		os.Exit(1)
	}

	out, err := os.Create("gopherexported.jpg")
	if err != nil {
		fmt.Printf("Error creating target file '%s'\n%v\n", out.Name(), err)
		os.Exit(1)
	}
	defer out.Close()

	out.Write(gopher)
	out.Sync()

}
