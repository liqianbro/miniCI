package templates

func MainTemplate() []byte {
	return []byte(`/*
{{ .Copyright }}
{{ if .Legal.Header }}{{ .Legal.Header }}{{ end }}
*/
package main

import "fmt"

func main() {
	fmt.Println("hello,world")
}
`)
}
