/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/wlbr/mule
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package {{.Package}}

import "encoding/base64"

// {{.Name}}Resource is a generated function returning the Resource as a []byte.
func {{.Name}}Resource() ([]byte, error) {
	var resource = "{{.Content}}"

	return base64.StdEncoding.DecodeString(resource)
}
