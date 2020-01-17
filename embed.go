/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/wlbr/templify
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package main

// embedTemplate is a generated function returning the template as a string.
// That string should be parsed by the functions of the golang's template package.
func embedTemplate() string {
	var tmpl = "/*\n" +
		" * CODE GENERATED AUTOMATICALLY WITH\n" +
		" *    github.com/wlbr/mule\n" +
		" * THIS FILE SHOULD NOT BE EDITED BY HAND\n" +
		" */\n" +
		"\n" +
		"package {{.Package}}\n" +
		"\n" +
		"import \"encoding/base64\"\n" +
		"\n" +
		"// {{.Name}}Resource is a generated function returning the Resource as a []byte.\n" +
		"func {{.Name}}Resource() ([]byte, error) {\n" +
		"\tvar resource = \"{{.Content}}\"\n" +
		"\n" +
		"\treturn base64.StdEncoding.DecodeString(resource)\n" +
		"}\n" +
		""
	return tmpl
}
