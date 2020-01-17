package main

//go:generate mule gopher.jpg

import (
	"fmt"
	"os"
)

func main() {
	// once you did a 'go generate' and then a 'go build' in this folder, you will have
	// an executable 'example' that has the image gopher.jpg included. (See line 3 above).
	// So if you run 'mulex' now, then it will create a new jpg in you filesystem
	// called 'gopherexported.jpg'.

	//Embedded template

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
