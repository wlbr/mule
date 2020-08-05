/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/wlbr/mule
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package {{.Package}}

	{{ if not .EmptyContent }} import "encoding/base64"
{{ end }}

// {{.Name}}Resource is a generated function returning the Resource as a []byte.
func {{.Name}}Resource() ([]byte, error) {
	var resource = "{{.Content}}"

	return {{ if .EmptyContent }} []byte(resource), nil{{ else }}base64.StdEncoding.DecodeString(resource){{ end }}
}
