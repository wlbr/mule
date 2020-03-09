package main

import (
	"fmt"
	"io/ioutil"
	"testing"
)

var SOURCETEMPLATE="embed.tpl"

func TestGeneratedTemplateFunktion(t *testing.T) {
	// This test is rather nonsense, as it tests a generated funktion
	// but it is needed as the guys from awesome-go want a higher test coverage.

	template := embedTemplate()

	if len(template) < 1 {
		fmt.Printf("Embedded, generated function 'embedTemplate()' does not return a full string\n")
		t.Fail()
	}

	sourcetemplatebytearray, err := ioutil.ReadFile(SOURCETEMPLATE)
	if err != nil {
		fmt.Printf("Could not read source template file '%s'.\n", SOURCETEMPLATE)
		t.Fail()
	}
	sourcetemplate:=string(sourcetemplatebytearray)
	if sourcetemplate != template {
		fmt.Printf("Source template file '%s' and result of generated template function differ.\n", SOURCETEMPLATE)
		t.Fail()
	}

}

